package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ShowEventHandler(ctx *gin.Context) {
	//check id
	id := ctx.Query("id")
	//if id is empty
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameters").Error())
		return
	}

	event := schemas.Event{}
	if err := db.First(&event, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "event not found")
		return
	}

	sendSuccess(ctx, "show event", event)
}
