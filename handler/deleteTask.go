package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteTaskHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	//Find task
	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("task with id: %s not found", id))
		return
	}

	//Delete task
	if err := db.Delete(&task, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete task with id: %s", id))
		return
	}

	sendSuccess(ctx, "task deleted", task)
}
