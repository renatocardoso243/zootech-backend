package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateEventHanlder(ctx *gin.Context) {
	request := CreateEventRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	event := schemas.Event{
		EventName: request.EventName,
		EventDate: request.EventDate,
	}

	if err := db.Create(&event).Error; err != nil {
		logger.Errorf("Error to create event: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating event on database")
		return
	}

	sendSuccess(ctx, "event created", event)
}
