package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateTaskHandler(ctx *gin.Context) {
	request := UpdateTaskRequest{}

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

	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "task not found")
		return
	}

	// Update task
	if request.EmployeeID != 0 {
		task.EmployeeID = request.EmployeeID
	}

	if request.TaskName != "" {
		task.TaskName = request.TaskName
	}

	if request.TaskTime != "" {
		task.TaskTime = request.TaskTime
	}

	if request.TaskDay != "" {
		task.TaskDay = request.TaskDay
	}

	if err := db.Save(&task).Error; err != nil {
		logger.Errorf("Error to update task: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating task on database")
		return
	}

	sendSuccess(ctx, "task updated", task)
}
