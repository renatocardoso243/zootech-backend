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
		// Animal routes
		v1.POST("/animals/create/", handler.CreateAnimalHandler)
		v1.GET("/animal", handler.ShowAnimalHandler)
		v1.DELETE("/animal", handler.DeleteAnimalHandler)
		v1.PUT("/animals/update/", handler.UpdateAnimalHandler)
		v1.GET("/animals", handler.ListAnimalsHandler)
		v1.GET("/animals-by-herd/:id", handler.ListAnimalsByHerdHandler)

		// Herd routes
		v1.POST("/herds/create/", handler.CreateHerdHandler)
		v1.GET("/herd", handler.ShowHerdHandler)
		v1.DELETE("/herd", handler.DeleteHerdHandler)
		v1.PUT("/herds/update/", handler.UpdateHerdHandler)
		v1.GET("/herds", handler.ListHerdsHandler)

		// Individual diet routes
		v1.POST("/diets/create/", handler.CreateIndividualDietHandler)
		v1.GET("/diet", handler.ShowIndividualDietHandler)
		v1.DELETE("/diet", handler.DeleteIndividualDietHandler)
		v1.PUT("/diets/update/", handler.UpdateIndividualDietHandler)
		v1.GET("/diets", handler.ListIndividualDietsHandler)

		// Weight control routes
		v1.POST("/weights/create/", handler.CreateWeightHandler)
		v1.DELETE("/weight", handler.DeleteWeightHandler)
		v1.PUT("/weights/update/", handler.UpdateWeightHandler)
		v1.GET("/weights", handler.ListWeightsHanlder)
	}
}