package routes

import (
	"github.com/OrdinarilySam/feddit/controllers"
	"github.com/OrdinarilySam/feddit/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.POST("/posts", middleware.RequireAuth, controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.DeletePost)
}
