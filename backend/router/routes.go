package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/auth"
	"github.com/shh4und/movie-tracker/config"
	"github.com/shh4und/movie-tracker/handler"
	"github.com/shh4und/movie-tracker/services"
)

func initRoutes(router *gin.Engine) {

	handler.InitHandler()
	services.InitServices()
	secret := []byte(config.Envs.JwtToken)

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register", handler.CreateUser) // create
		public.POST("/login", handler.LoginUser)     // authenticate
		// fetch a list of a title search
		public.GET("/titles/search", handler.GetTitle)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(auth.Authenticate(secret)) // using middleware to verify the authenticate session
	{
		// CRUD user routes
		protected.GET("/users/profile", handler.GetUserProfileByUsername)
		protected.PUT("/users/update", handler.UpdateUser)
		protected.DELETE("/users/delete", handler.DeleteUser)

		// User actions
		protected.POST("/rate", handler.AddRating)
		protected.DELETE("/rate", handler.RemoveRating)

		protected.POST("/comment", handler.AddComment)

		protected.POST("/favorite", handler.AddFavorite)
		protected.DELETE("/favorite", handler.RemoveFavorite)

		protected.POST("/watchlater", handler.AddWatchLater)

		protected.POST("/watched", handler.AddWatched)
	}

}
