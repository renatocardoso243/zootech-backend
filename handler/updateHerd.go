package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateHerdHandler(ctx *gin.Context) {
	request := UpdateHerdRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamenters").Error())
		return
	}
	
	herd := schemas.Herd{}
	if err := db.First(&herd, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "herd not found")
		return
	}

	// Update herd
	if request.HerdName != "" {
		herd.HerdName = request.HerdName
	}

	if request.Type != "" {
		herd.Type = request.Type
	}

	if request.Description != "" {
		herd.Description = request.Description
	}

	if err := db.Save(&herd).Error; err != nil {
		logger.Errorf("Error to update herd: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating herd on database")
		return
	}

	sendSuccess(ctx, "herd updated", herd)
}