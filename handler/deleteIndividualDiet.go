package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteIndividualDietHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	//Find herd
	individualDiet := schemas.IndividualDiet{}
	if err := db.First(&individualDiet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("diet with id: %s not found", id))
		return
	}

	//Delete herd
	if err := db.Delete(&individualDiet, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete diet with id: %s", id))
		return
	}

	sendSuccess(ctx, "diet deleted", individualDiet)
}