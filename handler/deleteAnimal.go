package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteAnimalHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}
	// Find Animal
	animal := schemas.Animal{}
	if err := db.First(&animal, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("animal with id: %s not found", id))
		return
	}

	// Delete Animal
	if err := db.Delete(&animal, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete animal with id: %s", id))
		return
	}
	sendSuccess(ctx, "animal deleted", animal)
}