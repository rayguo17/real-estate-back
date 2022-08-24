package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/goBack/routers/controllers"
)

func InitCORS(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))
	//router.Use(cors.Default())
	//router.Use(CORSMiddleware())
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		fmt.Println("hello? middleware?")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	InitCORS(r)
	v1 := r.Group("/api/v1/")
	//cors

	//
	{
		//global test
		site := v1.Group("site")
		{
			//for test
			site.GET("ping", controllers.Ping)
		}
		//track
		track := v1.Group("track")
		{
			track.GET("/users", controllers.GetUsers)
			track.GET("/user/:id", controllers.GetUser)
			track.POST("/user", controllers.CreateUser)
		}
		//LogBut
		metric := v1.Group("metric")
		{
			metric.GET("/logBtn", controllers.LogBtnHandler)
		}

	}
	return r
}
