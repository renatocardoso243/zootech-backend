package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/config"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListEmployeeTaskHandler(ctx *gin.Context) {
	// Get the employee id by the route
	employeeId := ctx.Param("id")
	if employeeId == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "pathParameters").Error())
		return
	}

	// var to store the tasks
	var tasks []schemas.Task

	// check tasks associated with the employee
	if err := config.DB.Where("employee_id = ?", employeeId).Find(&tasks).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error retrieving tasks for the specified employee")
		return
	}

	// Check is there is tasks for employee
	if len(tasks) == 0 {
		sendError(ctx, http.StatusNotFound, "no tasks found for the specified employee")
		return
	}

	// Return the tasks find
	sendSuccess(ctx, "list tasks successfully", tasks)
}
