package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateWeightHandler(ctx *gin.Context) {
	request := UpdateWeightRequest{}

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

	// Update weight register
	weight := schemas.Weight{}
	if err := db.First(&weight, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "weight register not found")
		return
	}

	if request.Data != "" {
		weight.Data = request.Data
	}

	if request.TotalWeight != "" {
		weight.TotalWeight = request.TotalWeight
	}

	if request.Gain != "" {
		weight.Gain = request.Gain
	}

	if request.Loss != "" {
		weight.Loss = request.Loss
	}

	if request.AnimalID != 0 {
		weight.AnimalID = request.AnimalID
	}

	// Save weight register
	if err := db.Save(&weight).Error; err != nil {
		logger.Errorf("Error to update weight register: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating weight register on database")
		return
	}

	sendSuccess(ctx, "weight register updated", weight)

}