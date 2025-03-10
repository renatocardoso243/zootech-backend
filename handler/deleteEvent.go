package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteEventHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameters").Error())
		return
	}

	// Find id
	event := schemas.Event{}
	if err := db.First(&event, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("event with id: %s not found", id))
		return
	}

	// Delete event
	if err := db.Delete(&event, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete event with id: %s", id))
		return
	}

	sendSuccess(ctx, "event deleted", event)
}
