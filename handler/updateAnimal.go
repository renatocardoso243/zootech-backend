package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateAnimalHandler(ctx *gin.Context) {
	request := UpdateAnimalRequest{}

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

	animal := schemas.Animal{}
	if err := db.First(&animal, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "animal not found")
		return
	}

	// Update animal
	if request.Name != "" {
		animal.Name = request.Name
	}

	if request.Tag != "" {
		animal.Tag = request.Tag
	}

	if request.Race != "" {
		animal.Race = request.Race
	}

	if request.Weight != 0 {
		animal.Weight = request.Weight
	}

	if request.Sex != "" {
		animal.Sex = request.Sex
	}

	if request.Birthdate != "" {
		animal.Birthdate = request.Birthdate
	}

	if request.Vaccinated != "" {	
		animal.Vaccinated = request.Vaccinated
	}

	if request.Health != "" {
		animal.Health = request.Health
	}

	if request.HerdID != 0 {	
		animal.HerdID = request.HerdID
	}	

	if request.Observation != "" {	
		animal.Observation = request.Observation
	}

	// Save animal
	if err := db.Save(&animal).Error; err != nil {
		logger.Errorf("Error to update animal: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating animal on database")
		return
	}

	sendSuccess(ctx, "animal updated", animal)
}