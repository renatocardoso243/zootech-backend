package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ShowAnimalHandler(ctx *gin.Context) {
	// Check id
	id := ctx.Query("id")
	// If id is empty
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	animal := schemas.Animal{}
	if err := db.First(&animal, id).Error;err != nil {
		sendError(ctx, http.StatusNotFound, "animal not found")
		return
	}

	sendSuccess(ctx, "show animal", animal)
}