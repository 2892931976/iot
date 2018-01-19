//
//     Schemes: http, https
//     Host: localhost:3333
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: 周庆<admin@mojotv.cn> http://www.mojotv.cn
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/mojoGin/models"
	"github.com/mojocn/turbo-iot/config"
	"github.com/mojocn/turbo-iot/controllers"
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
	router.POST("/iot/v1/login", controllers.AuthPost)

	iotV1 := router.Group("/iot/v1")
	iotV1.Use(middlewares.JwtTokenCheck())
	{
		iotV1.GET("/device", controllers.DeviceIndex)
		iotV1.POST("/device", controllers.DeviceAdd)
		iotV1.GET("/device/:dno", controllers.DeviceInfo)
		iotV1.PUT("/device", controllers.DeviceUpdate)
		iotV1.DELETE("/device/:dno", controllers.DeviceDelete)

		iotV1.POST("/command", controllers.CommandAdd)
	}

	router.Run(config.AppPort)
}
