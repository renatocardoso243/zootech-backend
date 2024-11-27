package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateIndividualDietHandler(ctx *gin.Context) {
	request := UpdateIndividualDietRequest{}

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

	individualDiet := schemas.IndividualDiet{}
	if err := db.First(&individualDiet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "diet not found")
		return
	}

	// Update diet
	if request.DietName != "" {
		individualDiet.DietName = request.DietName 
	}

	if request.Type != "" {
		individualDiet.Type = request.Type
	}

	if request.StartDate != "" {
		individualDiet.StartDate = request.StartDate
	}

	if request.EndDate != "" {
		individualDiet.EndDate = request.EndDate
	}

	if request.Description != "" {
		individualDiet.Description = request.Description
	}

	// Save diet
	if err := db.Save(&individualDiet).Error; err != nil {
		logger.Errorf("Error to update diet: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating diet on database")
		return
	}

	sendSuccess(ctx, "diet updated", individualDiet)
}