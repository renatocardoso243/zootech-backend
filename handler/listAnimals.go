package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListAnimalsHandler(ctx *gin.Context) {
	animal := []schemas.Animal{}

	if err := db.Find(&animal).Error;err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list animals")
		return
	}

	sendSuccess(ctx, "list animals", animal)
}