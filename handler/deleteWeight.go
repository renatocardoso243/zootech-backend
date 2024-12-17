package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteWeightHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	// Find weight register
	weight := schemas.Weight{}
	if err := db.First(&weight, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("weight register with id: %s not found", id))
		return
	}

	// Delete register
	if err := db.Delete(&weight, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete weight register with id: %s", id))
		return
	}

	sendSuccess(ctx, "weight register deleted", weight)
}