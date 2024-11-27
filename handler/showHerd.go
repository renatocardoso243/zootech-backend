package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/config"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ShowHerdHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}
	herd := schemas.Herd{}
	if err := config.DB.Preload("Animals").First(&herd, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "herd not found")
		return
	}
	sendSuccess(ctx,"show herd", herd)
}

