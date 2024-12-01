package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteGroupDietHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	//Find herd
	groupDiet := schemas.GroupDiet{}
	if err := db.First(&groupDiet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("diet with id: %s not found", id))
		return
	}

	//Delete herd
	if err := db.Delete(&groupDiet, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete diet with id: %s", id))
		return
	}

	sendSuccess(ctx, "diet deleted", groupDiet)
}