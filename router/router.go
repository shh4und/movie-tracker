package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	initRoutes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
