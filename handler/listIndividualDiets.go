package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListIndividualDietsHandler(ctx *gin.Context) {
	individualDiet := []schemas.IndividualDiet{}

	if err:= db.Find(&individualDiet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list diets")
		return
	}

	sendSuccess(ctx, "list diets", individualDiet)
}