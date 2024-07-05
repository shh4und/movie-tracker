package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/auth"
	"github.com/shh4und/movie-tracker/config"
	"github.com/shh4und/movie-tracker/handler"
)

func initRoutes(router *gin.Engine) {

	handler.InitHandler()
	secret := []byte(config.Envs.JwtToken)
	titles := router.Group("/api/v1/search")
	{
		// fetch a list of a title search
		titles.GET("", handler.GetSearch)
	}

	// public routes
	router.POST("/api/v1/register", handler.CreateUser) // create
	router.POST("/api/v1/login", handler.LoginUser)     // authenticate

	usersProtected := router.Group("/api/v1/users")
	usersProtected.Use(auth.Authenticate(secret))
	{

		// CRUD user routes

		// users.GET("/profile/id", handler.GetUserProfileByID)
		usersProtected.GET("/profile", handler.GetUserProfileByUsername)
		usersProtected.PUT("/update", handler.UpdateUser)
		usersProtected.DELETE("/delete", handler.DeleteUser)

	}

}
