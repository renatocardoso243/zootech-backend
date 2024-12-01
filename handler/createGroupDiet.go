package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateGroupDietHandler(ctx *gin.Context) {
	request := CreateGroupDietRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation erro: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	groupDiet := schemas.GroupDiet{
		DietName:    request.DietName,
		Type:        request.Type,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
		HerdID:    	 request.HerdID,
	}
	
	if err := db.Create(&groupDiet).Error; err != nil {
		logger.Errorf("Error to create individual diet: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating individual diet on database")
		return
	}

	sendSuccess(ctx, "diet created", groupDiet)
}