package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ShowTaskHandler(ctx *gin.Context) {
	// Check id
	id := ctx.Query("id")
	// If id is empty
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "task not found")
		return
	}

	sendSuccess(ctx, "show task", task)
}
