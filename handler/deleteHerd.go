package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func DeleteHerdHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	//Find herd
	herd := schemas.Herd{}
	if err := db.First(&herd, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("herd with id: %s not found", id))
		return
	}

	//Delete herd
	if err := db.Delete(&herd, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error to delete herd with id: %s", id))
		return
	}

	sendSuccess(ctx, "herd deleted", herd)
}