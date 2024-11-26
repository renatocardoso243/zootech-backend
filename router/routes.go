package router

import (
	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/handler"
)

func initializeRoutes(router *gin.Engine) {

	// Initialize Hanlder
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		// Animais
		v1.POST("/animal", handler.CreateAnimalHandler)
		v1.GET("/animal", handler.ShowAnimalHandler)
		v1.DELETE("/amimal", handler.DeleteAnimalHandler)
		v1.PUT("/animal", handler.UpdateAnimalHandler)
		v1.GET("/animals", handler.ListAnimalsHandler)
		// Rebanho
		v1.POST("/herd", handler.CreateHerdHandler)
		v1.GET("/herd", handler.ShowHerdHandler)
	}
}