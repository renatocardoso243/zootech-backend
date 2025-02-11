package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteEmployeeHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	// Find Employee
	employee := schemas.Employee{}
	if err := db.First(&employee, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("employee with id: %s not found", id))
		return
	}

	// Delete Employee
	if err := db.Delete(&employee, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete employee with id: %s", id))
		return
	}
	sendSuccess(ctx, "employee deleted", employee)
}
