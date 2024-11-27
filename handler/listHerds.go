package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListHerdsHandler(ctx *gin.Context) {
	herd := []schemas.Herd{}

	if err := db.Find(&herd).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list herds")
		return
	}

	sendSuccess(ctx, "list herds", herd)
}