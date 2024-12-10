package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateIndividualDietHandler(ctx *gin.Context) {
	request := CreateDietRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation erro: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	individualDiet := schemas.Diet{
		DietName:    		request.DietName,
		AnimalType:  		request.AnimalType,
		Objective:    		request.Objective,
		Ingredients: 		request.Ingredients,
		NutritionalInfo: 	request.NutritionalInfo,
	}
	
	if err := db.Create(&individualDiet).Error; err != nil {
		logger.Errorf("Error to create individual diet: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating individual diet on database")
		return
	}

	sendSuccess(ctx, "diet created", individualDiet)
}