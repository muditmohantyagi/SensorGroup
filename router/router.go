package router

import (
	"github.com/gin-contrib/cors"
	"test.com/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//add swager rout
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Max"}
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("POST", "GET", "PUT", "DELETE", "UPDATE", "OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))
	/**Allow origin CORS setting end:**/

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})
	/*-------------Routing started---------------*/
	user := r.Group("api/")
	{
		var GroupController = new(controllers.GroupController)
		user.POST("/group_sensor_generation", GroupController.GroupSensorGeneration)
		user.POST("/data_generation", GroupController.DataGeneration)

		user.POST("/group/species/:groupName", GroupController.GroupSpecies)
		user.POST("/data_list", GroupController.ShowAllData)
		user.POST("/group/transparency/average/:groupName", GroupController.Transparency)
		user.POST("/group/temperature/average/:groupName", GroupController.Temperature)
		user.POST("/sensor/average/temperature/:codeName", GroupController.SenserTemperature)
		user.POST("/group/species/top/:groupName/:N", GroupController.TopSpecies)

		user.POST("/region/temperature/min", GroupController.MinTempReason)
		user.POST("/region/temperature/max", GroupController.MaxTempReason)

	}
	return r

}
