package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListEmployeeHandler(ctx *gin.Context) {
	employee := []schemas.Employee{}

	if err := db.Find(&employee).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list employees")
		return
	}
	sendSuccess(ctx, "list employees", employee)
}
