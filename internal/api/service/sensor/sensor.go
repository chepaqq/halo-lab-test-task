package sensor

import (
	"github.com/chepaqq99/halo-lab-test-task/internal/models"
)

type sensorRepository interface {
	CreateSensor(index, groupID, coordinatesID, dataOutputRate int) (int, error)
	CreateCoordinates(coordinates *models.Coordinates) (int, error)
	CreateFish(fishName string) (int, error)
	UpdateSensor(index, groupID, transparency int, temperature float64) (int, error)
}

type SensorService struct {
	repo sensorRepository
}

func NewSensorService(repo sensorRepository) *SensorService {
	return &SensorService{repo: repo}
}

func (s *SensorService) CreateSensor(index, groupID, coordinatesID, dataOutputRate int) (int, error) {
	return s.repo.CreateSensor(index, groupID, coordinatesID, dataOutputRate)
}

func (s *SensorService) UpdateSensor(index, groupID, transparency int, temperature float64) (int, error) {
	return s.repo.UpdateSensor(index, groupID, transparency, temperature)
}

func (s *SensorService) CreateCoordinates(coordinates *models.Coordinates) (int, error) {
	return s.repo.CreateCoordinates(coordinates)
}

func (s *SensorService) CreateFish(fishName string) (int, error) {
	return s.repo.CreateFish(fishName)
}
