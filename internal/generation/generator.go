package generation

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	groupRepository "github.com/chepaqq99/halo-lab-test-task/internal/api/repository/group"
	sensorRepository "github.com/chepaqq99/halo-lab-test-task/internal/api/repository/sensor"
	groupService "github.com/chepaqq99/halo-lab-test-task/internal/api/service/group"
	sensorService "github.com/chepaqq99/halo-lab-test-task/internal/api/service/sensor"

	"github.com/chepaqq99/halo-lab-test-task/internal/models"
	"github.com/chepaqq99/halo-lab-test-task/pkg/db"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

// TODO: Update sensor table and log into sensor_history

var greekLettersName = []string{
	"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta",
	"theta", "iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi",
	"rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega",
}

// InitGeneration - .
func InitGeneration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfgDB := db.Config{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("POSTGRES_HOST"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	db, err := db.ConnectPostgres(cfgDB)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	groupRepository := groupRepository.NewGroupDB(db, nil)
	groupService := groupService.NewGroupService(groupRepository)

	sensorRepository := sensorRepository.NewSensorDB(db)
	sensorService := sensorService.NewSensorService(sensorRepository)

	// one-time "kickoff" phase
	sensors := createGroupsAndSensors(groupService, sensorService, 5)

	fishSpecies := parseFishSpecies()
	createFishes(sensorService, fishSpecies)

	// regularly repeated phase
	sem := make(chan struct{}, 5) // semaphore to control concurrent database operations
	for i := range sensors {
		go generateData(&sensors[i], &sensors, &fishSpecies, sensorService, sem)
	}
}

func createGroupsAndSensors(groupService *groupService.GroupService, sensorService *sensorService.SensorService, numSensorsPerGroup int) []models.Sensor {
	var sensors []models.Sensor
	for groupID, groupName := range greekLettersName {
		_, _ = groupService.CreateGroup(groupName)
		lastIndex, _ := groupService.GetLastIDInGroup(groupID + 1)
		for i := 1; i <= numSensorsPerGroup; i++ {
			coordinates := generateCoordinates()
			coordinatesID, _ := sensorService.CreateCoordinates(&coordinates)
			sensorIndex := lastIndex + i
			dataOutputRate := rand.Intn(5) + 1
			_, err := sensorService.CreateSensor(sensorIndex, groupID+1, coordinatesID, dataOutputRate)
			if err != nil {
				log.Fatal(err)
			}
			sensors = append(
				sensors,
				models.Sensor{
					Index:          sensorIndex + i,
					GroupID:        groupID + 1,
					DataOutputRate: time.Duration(dataOutputRate) * time.Second,
					Coordinates: models.Coordinates{
						X: coordinates.X,
						Y: coordinates.Y,
						Z: coordinates.Z,
					},
				},
			)
		}
	}
	return sensors
}

func createFishes(sensorService *sensorService.SensorService, fishSpecies []string) {
	for _, fishSpecie := range fishSpecies {
		_, _ = sensorService.CreateFish(fishSpecie)
	}
}

func generateCoordinates() models.Coordinates {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	minCoord := -100
	maxCoord := 100

	x := r.Intn(maxCoord-minCoord+1) + minCoord
	y := r.Intn(maxCoord-minCoord+1) + minCoord
	z := -r.Intn(maxCoord-minCoord+1) - 1

	return models.Coordinates{X: x, Y: y, Z: z}
}

func parseFishSpecies() []string {
	c := colly.NewCollector(
		colly.AllowedDomains("oceana.org"),
	)
	var names []string
	c.OnHTML(".tb-grid", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			name := strings.TrimSpace(el.Text)
			if name != "" {
				names = append(names, name)
			}
		})
	})
	err := c.Visit("https://oceana.org/ocean-fishes/")
	if err != nil {
		log.Fatal(err)
	}
	return names
}

func generateData(sensor *models.Sensor, allSensors *[]models.Sensor, fishSpecies *[]string, sensorService *sensorService.SensorService, sem chan struct{}) {
	ticker := time.NewTicker(sensor.DataOutputRate)
	defer ticker.Stop()

	for range ticker.C {
		sensor.Transparency = rand.Intn(100)

		depth := sensor.Coordinates.Z
		sensor.Temperature = generateTemperature(depth)

		maxTransparencyDiff := 10 // Maximum allowed difference in transparency for nearby sensors
		// Ensure transparency for nearby sensors does not differ too much
		for _, s := range *allSensors {
			if s != *sensor && isNearby(sensor.Coordinates, s.Coordinates) {
				if abs(sensor.Transparency-s.Transparency) > maxTransparencyDiff {
					if sensor.Transparency > s.Transparency {
						sensor.Transparency = s.Transparency + maxTransparencyDiff
					} else {
						sensor.Transparency = s.Transparency - maxTransparencyDiff
					}
				}
			}
		}

		// Acquire semaphore before performing the database update
		sem <- struct{}{}
		_, err := sensorService.UpdateSensor(sensor.Index, sensor.GroupID, sensor.Transparency, sensor.Temperature)
		<-sem // Release semaphore after the database update is complete

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Sensor %d - Temperature: %.2f°C, Transparency: %d%%\n", sensor.Index, sensor.Temperature, sensor.Transparency)
	}
}

func generateTemperature(depth int) float64 {
	return float64(depth) / 100.0 * 30.0
}

func isNearby(c1, c2 models.Coordinates) bool {
	dx := math.Abs(float64(c1.X - c2.X))
	dy := math.Abs(float64(c1.Y - c2.Y))
	dz := math.Abs(float64(c1.Z - c2.Z))
	return dx <= 10 && dy <= 10 && dz <= 10
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func generateDetectedSpecies(fishSpeciesNames []string) map[string]int {
	numSpecies := rand.Intn(6)
	detectedSpecies := make(map[string]int)

	for i := 0; i < numSpecies; i++ {
		speciesIndex := rand.Intn(len(fishSpeciesNames))
		species := fishSpeciesNames[speciesIndex]
		count := rand.Intn(10) + 1
		detectedSpecies[species] = count
	}

	return detectedSpecies
}
