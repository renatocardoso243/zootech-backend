package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {

	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code, 
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler:  %s successfull", op), 
		"data": data,})
}

type ErrorResponse struct {
	Message 	string `json:"message"`
	ErrorCode 	string `json:"errorCode"`
}


// Animal
type CreateAnimalResponse struct {
	Message string `json:"message"`
	Data 	schemas.Animal `json:"data"`
}
type DeleteAnimalResponse struct {
	Message string `json:"message"`
	Data 	schemas.Animal `json:"data"`
}
type ShowAnimalResponse struct {
	Message string `json:"message"`
	Data 	schemas.Animal `json:"data"`
}
type ListAnimalsResponse struct {
	Message string `json:"message"`
	Data 	[]schemas.Animal `json:"data"`
}

type UpdateAnimalResponse struct {
	Message string `json:"message"`
	Data 	schemas.Animal `json:"data"`
}

// Herd
type CreateHerdResponse struct {
	Message string `json:"message"`
	Data 	schemas.Herd `json:"data"`
}

type DeleteHerdResponse struct {
	Message string `json:"message"`
	Data 	schemas.Herd `json:"data"`
}

type ShowHerdResponse struct {
	Message string `json:"message"`
	Data 	schemas.Herd `json:"data"`
}

type ListHerdsResponse struct {
	Message string `json:"message"`
	Data 	[]schemas.Herd `json:"data"`
}

type UpdateHerdResponse struct {
	Message string `json:"message"`
	Data 	schemas.Herd `json:"data"`
}

//Diet
type CreateIndividualDietResponse struct {
	Message string `json:"message"`
	Data 	schemas.Diet `json:"data"`
}

type DeleteIndividualDietResponse struct {
	Message string `json:"message"`
	Data 	schemas.Diet `json:"data"`
}

type ShowIndividualDietResponse struct {
	Message string `json:"message"`
	Data 	schemas.Diet `json:"data"`
}

type ListIndividualDietsResponse struct {
	Message string `json:"message"`
	Data 	[]schemas.Diet `json:"data"`
}

type UpdateIndividualDietResponse struct {
	Message string `json:"message"`
	Data 	schemas.Diet `json:"data"`
}

