package router

import (
	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/middleware"
)

func Initialize() {
	// Initialize Router with CORS middleware
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	initializeRoutes(router)
	
	router.Run(":8080")
}