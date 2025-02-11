package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ShowEmployeeHandler(ctx *gin.Context) {
	// Check id
	id := ctx.Query("id")
	// If id is empty
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	employee := schemas.Employee{}
	if err := db.First(&employee, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "employee not found")
		return
	}

	sendSuccess(ctx, "show employee", employee)
}
