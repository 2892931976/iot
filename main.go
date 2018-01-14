package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/mojoGin/models"
	"github.com/mojocn/turbo-iot/controllers"
	"time"
)

func main() {
	//init Redis
	defer models.DB.Close()

	router := gin.Default()
	router.Use(gin.Logger())

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorize"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 24 * 365 * time.Hour,
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.

	router.GET("/", controllers.AuthPost)

	v1 := router.Group("/v1")
	{
		v1.POST("/login", controllers.AuthPost)
	}

	router.Run(":3333")
}
