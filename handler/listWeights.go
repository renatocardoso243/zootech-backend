package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func ListWeightsHanlder(ctx *gin.Context) {
	weight := []schemas.Weight{}

	if err := db.Find(&weight).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to list weight register")
		return
	}

	sendSuccess(ctx, "list weight register", weight)
}