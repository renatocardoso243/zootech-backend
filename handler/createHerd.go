package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateHerdHandler(ctx *gin.Context) {
	request := CreateHerdRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation erro: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	herd := schemas.Herd{
		HerdName:    request.HerdName,
		Type:        request.Type,
		Description: request.Description,
	}

	if err := db.Create(&herd).Error; err !=  nil {
		logger.Errorf("Error to create herd: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating herd on database")
		return
	}

	sendSuccess(ctx, "herd created", herd)
}