package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateAnimalHandler(ctx *gin.Context) {
	request := CreateAnimalRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// create animal
	animal := schemas.Animal{
		Name:        request.Name,
		Tag:         request.Tag,
		Race:        request.Race,
		Weight:      request.Weight,
		Sex:         request.Sex,
		Birthdate:   request.Birthdate,
		Vaccinated:  request.Vaccinated,
		Health:		 request.Health,
		HerdID:      request.HerdID,
		Observation: request.Observation,
	}

	if err := db.Create(&animal).Error; err != nil {
		logger.Errorf("Error to create animal: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating animal on database")
		return
	}

	sendSuccess(ctx, "animal created", animal)
	
}