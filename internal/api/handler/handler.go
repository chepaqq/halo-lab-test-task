package handler

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	router := gin.New()
	group := router.Group("/group")
	{
		groupName := group.Group("/:groupName")
		{
			groupName.GET("/transparency/average")
			groupName.GET("/temperature/average")
			groupName.GET("/species/")
			groupName.GET("/species/top/:N")
		}
	}
	router.GET("/sensor/:codeName/temperature/average")
	router.GET("/region/temperature/max")
	router.GET("/region/temperature/min")

	return router
}
