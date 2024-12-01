package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateGroupDietHandler(ctx *gin.Context) {
	request := UpdateGroupDietRequest{}

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

	groupDiet := schemas.GroupDiet{}
	if err := db.First(&groupDiet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "diet not found")
		return
	}

	// Update diet
	if request.DietName != "" {
		groupDiet.DietName = request.DietName 
	}

	if request.Type != "" {
		groupDiet.Type = request.Type
	}

	if request.StartDate != "" {
		groupDiet.StartDate = request.StartDate
	}

	if request.EndDate != "" {
		groupDiet.EndDate = request.EndDate
	}

	if request.Description != "" {
		groupDiet.Description = request.Description
	}

	// Save diet
	if err := db.Save(&groupDiet).Error; err != nil {
		logger.Errorf("Error to update diet: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating diet on database")
		return
	}

	sendSuccess(ctx, "diet updated", groupDiet)
}