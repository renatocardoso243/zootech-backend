package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateIndividualDietHandler(ctx *gin.Context) {
	request := CreateIndividualDietRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation erro: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	individualDiet := schemas.IndividualDiet{
		DietName:    request.DietName,
		Type:        request.Type,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
		AnimalID:    request.AnimalID,
	}
	
	if err := db.Create(&individualDiet).Error; err != nil {
		logger.Errorf("Error to create individual diet: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating individual diet on database")
		return
	}

	sendSuccess(ctx, "diet created", individualDiet)
}