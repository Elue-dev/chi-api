package main

import (
	"github.com/elue-dev/gin-api/controllers"
	"github.com/elue-dev/gin-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectAndMigrateDB()
}

func main() {	
	r := gin.Default()

	r.POST("/posts", controllers.AddPost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run() 
}