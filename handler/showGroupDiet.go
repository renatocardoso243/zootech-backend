package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ShowGroupDietHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}

	groupDiet := schemas.GroupDiet{}
	if err := db.First(&groupDiet, id).Error;err != nil {
		sendError(ctx, http.StatusNotFound, "diet not found")
		return
	}
	sendSuccess(ctx, "show diet", groupDiet)
}