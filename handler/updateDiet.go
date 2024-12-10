package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateIndividualDietHandler(ctx *gin.Context) {
	request := UpdateDietRequest{}

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

	diet := schemas.Diet{}
	if err := db.First(&diet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "diet not found")
		return
	}

	// Update diet
	if request.DietName != "" {
		diet.DietName = request.DietName 
	}

	if request.AnimalType != "" {
		diet.AnimalType = request.AnimalType
	}

	if request.Objective != "" {
		diet.Objective = request.Objective
	}

	if request.Ingredients != "" {
		diet.Ingredients = request.Ingredients
	}

	if request.NutritionalInfo != "" {
		diet.NutritionalInfo = request.NutritionalInfo
	}

	// Save diet
	if err := db.Save(&diet).Error; err != nil {
		logger.Errorf("Error to update diet: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating diet on database")
		return
	}

	sendSuccess(ctx, "diet updated", diet)
}