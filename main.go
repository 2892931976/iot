package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/mojoGin/models"
	"github.com/mojocn/turbo-iot/config"
	"github.com/mojocn/turbo-iot/handlers"
	"github.com/mojocn/turbo-iot/middlewares"
	"time"
)

func main() {
	//init Redis
	defer models.DB.Close()

	router := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Origin", "Authorize", "Authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 24 * 365 * time.Hour,
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.POST("/iot/v1/login", handlers.AuthPost)

	iotV1 := router.Group("/iot/v1")
	iotV1.Use(middlewares.JwtTokenCheck())
	{
		iotV1.GET("/device", handlers.DeviceIndex)
		iotV1.POST("/device", handlers.DeviceAdd)
		iotV1.GET("/device/:dno", handlers.DeviceInfo)
		iotV1.PUT("/device", handlers.DeviceUpdate)
		iotV1.DELETE("/device/:dno", handlers.DeviceDelete)

		iotV1.POST("/command", handlers.CommandAdd)
	}

	router.Run(config.AppPort)
}
