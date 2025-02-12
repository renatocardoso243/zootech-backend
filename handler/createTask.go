package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateTaskHandler(ctx *gin.Context) {
	request := CreateTaskRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation erro: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	task := schemas.Task{
		EmployeeID:     request.EmployeeID,
		TaskName:       request.TaskName,
		TaskDate:       request.TaskDate,
		TaskTime:       request.TaskTime,
		TaskConclusion: request.TaskConclusion,
	}

	// Save task in database
	if err := db.Create(&task).Error; err != nil {
		logger.Errorf("Error to create task: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating task on database")
		return
	}

	sendSuccess(ctx, "task created", task)
}
