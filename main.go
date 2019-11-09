package main

import (
	"scg-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	scgGroup := router.Group("scg")
	{
		scg := new(controllers.SCGController)
		scgGroup.GET("/findXYZ", scg.FindXYZ)
		scgGroup.GET("/findPlace", scg.FindPlace)
	}

	router.Run(":3000")
}
