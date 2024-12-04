package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/config"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListAnimalsByHerdHandler(ctx *gin.Context) {
	// Get the herd id by the route
	herdId := ctx.Param("id")
	if herdId == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "pathParameters").Error())
		return
	}

	// var to store the animals
	var animals []schemas.Animal

	// Check animals associated with the herd
	if err := config.DB.Where("herd_id = ?", herdId).Find(&animals).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error retrieving animals for the specified herd")
		return
	}

	// Check is there is animals in the herd
	if len(animals) == 0 {
		sendError(ctx, http.StatusNotFound, "no animals found for the specified herd")
		return
	}

	// Return the animals find
	sendSuccess(ctx, "list animals successfully", animals)
}
