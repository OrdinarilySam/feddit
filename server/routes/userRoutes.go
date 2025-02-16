package routes

import (
	"github.com/OrdinarilySam/feddit/controllers"
	"github.com/OrdinarilySam/feddit/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
}
