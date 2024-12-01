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
		v1.POST("/animal", handler.CreateAnimalHandler)
		v1.GET("/animal", handler.ShowAnimalHandler)
		v1.DELETE("/animal", handler.DeleteAnimalHandler)
		v1.PUT("/animal", handler.UpdateAnimalHandler)
		v1.GET("/animals", handler.ListAnimalsHandler)
		// Herd routes
		v1.POST("/herd", handler.CreateHerdHandler)
		v1.GET("/herd", handler.ShowHerdHandler)
		v1.DELETE("/herd", handler.DeleteHerdHandler)
		v1.PUT("/herd", handler.UpdateHerdHandler)
		v1.GET("/herds", handler.ListHerdsHandler)
		// Individual diet routes
		v1.POST("/diet", handler.CreateIndividualDietHandler)
		v1.GET("/diet", handler.ShowIndividualDietHandler)
		v1.DELETE("/diet", handler.DeleteIndividualDietHandler)
		v1.PUT("/diet", handler.UpdateIndividualDietHandler)
		v1.GET("/diets", handler.ListIndividualDietsHandler)
		// Group diets routes
		v1.POST("/group-diet", handler.CreateGroupDietHandler)
		v1.GET("/group-diet", handler.ShowGroupDietHandler)
		v1.DELETE("/group-diet", handler.DeleteGroupDietHandler)
		v1.PUT("/group-diet", handler.UpdateGroupDietHandler)
		v1.GET("/group-diets", handler.ListGroupDietsHandler)
	}
}