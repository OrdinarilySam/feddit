package main

import (
	"github.com/OrdinarilySam/feddit/initializers"
	"github.com/OrdinarilySam/feddit/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.UserRoutes(r)
	routes.PostRoutes(r)

	r.Run()
}
