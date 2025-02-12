package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListTasksHandler(ctx *gin.Context) {
	task := []schemas.Task{}

	if err := db.Find(&task).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list tasks")
		return
	}

	sendSuccess(ctx, "list tasks", task)
}
