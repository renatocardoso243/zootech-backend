package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateEventHandler(ctx *gin.Context) {
	request := UpdateEventRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	event := schemas.Event{}
	if err := db.First(&event, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "event not found")
		return
	}

	// Update event
	if request.EventName != "" {
		event.EventName = request.EventName
	}

	if request.EventDescription != "" {
		event.EventDescription = request.EventDescription
	}

	if request.EventDate != "" {
		event.EventDate = request.EventDate
	}

	if err := db.Save(&event).Error; err != nil {
		logger.Errorf("Error to update event: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating event on database")
		return
	}

	sendSuccess(ctx, "event updated", event)

}
