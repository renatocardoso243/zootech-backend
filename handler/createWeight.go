package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateWeightHandler(ctx *gin.Context) {
	request := CreateWeightRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	weight := schemas.Weight{
		Data: 		request.Data,
		TotalWeight: request.TotalWeight,
		Gain: 		request.Gain,
		Loss: 		request.Loss,
		AnimalID: 	request.AnimalID,
	}

	if err := db.Create(&weight).Error; err != nil {
		logger.Errorf("Error to register weight: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error register weight on database")
		return
	}

	sendSuccess(ctx, "weight register", weight)
}