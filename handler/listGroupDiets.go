package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListGroupDietsHandler(ctx *gin.Context) {
	groupDiet := []schemas.GroupDiet{}

	if err:= db.Find(&groupDiet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list diets")
		return
	}

	sendSuccess(ctx, "list diets", groupDiet)
}