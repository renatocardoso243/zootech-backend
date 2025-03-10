package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListEventHandler(ctx *gin.Context) {
	event := []schemas.Event{}

	if err := db.Find(&event).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list events")
		return
	}

	sendSuccess(ctx, "list events", event)
}
