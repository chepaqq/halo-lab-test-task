package handler

import (
	"log"
	"os"

	groupHandler "github.com/chepaqq99/halo-lab-test-task/internal/api/handler/group"
	groupRepository "github.com/chepaqq99/halo-lab-test-task/internal/api/repository/group"
	groupService "github.com/chepaqq99/halo-lab-test-task/internal/api/service/group"
	"github.com/chepaqq99/halo-lab-test-task/pkg/db"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// InitRoutes - .
func InitRoutes() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := db.Config{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("POSTGRES_HOST"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}

	db, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	groupRepository := groupRepository.NewGroupDB(db)
	groupService := groupService.NewGroupService(groupRepository)
	groupHandler := groupHandler.NewGroupHandler(groupService)

	router := gin.New()
	group := router.Group("/group")
	{
		groupName := group.Group("/:groupName")
		{
			groupName.GET("/transparency/average", groupHandler.GetAverageTransparency)
			groupName.GET("/temperature/average", groupHandler.GetAverageTemperature)
			groupName.GET("/species", groupHandler.GetListOfSpecies)
		}
	}

	return router
}
