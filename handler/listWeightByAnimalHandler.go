package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/config"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListWeightByAnimalHandler(ctx *gin.Context) {
	// Get the animal id by the route
	animalId := ctx.Param("id")
	if animalId == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "pathParameters").Error())
		return
	}

	// var to store the weights
	var weights []schemas.Weight

	// Check weights associated with the animal
	if err := config.DB.Where("animal_id = ?", animalId).Find(&weights).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error retrieving weights for the specified animal")
		return
	}

	// Check if exist register of weight for the animal
	if len(weights) == 0 {
		sendError(ctx, http.StatusNotFound, "no weights found for the specified animal")
		return
	}

	// Return the registers
	sendSuccess(ctx, "list registers successfully", weights)
}