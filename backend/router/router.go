package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {

	// Initialize router (gin.Engine) with the default configs
	router := gin.Default()
	// Configure trusted proxies
	router.SetTrustedProxies([]string{"127.0.0.1"}) // Confia apenas em localhost
	// Serve static files from the frontend/static directory
	router.Static("/static", "../frontend/static")

	// Serve the main HTML files directly from frontend/static
	router.StaticFile("/index", "../frontend/static/index.html")
	router.StaticFile("/register", "../frontend/static/register.html")
	router.StaticFile("/login", "../frontend/static/login.html")

	// Initialize routes
	initRoutes(router)

	// Run server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
