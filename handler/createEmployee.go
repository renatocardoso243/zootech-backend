package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func CreateEmployeeHandler(ctx *gin.Context) {
	request := CreateEmployeeRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// create employee
	employee := schemas.Employee{
		EmployeeId:      request.EmployeeId,
		FullName:        request.FullName,
		Birthdate:       request.Birthdate,
		Genre:           request.Genre,
		CivilStatus:     request.CivilStatus,
		Nacionality:     request.Nacionality,
		TaxIdNumber:     request.TaxIdNumber,
		IdentityCard:    request.IdentityCard,
		Address:         request.Address,
		Phone:           request.Phone,
		Email:           request.Email,
		Role:            request.Role,
		AdmissionDate:   request.AdmissionDate,
		Departament:     request.Departament,
		ResignationDate: request.ResignationDate,
		WorkRegiment:    request.WorkRegiment,
		Salary:          request.Salary,
		BankData:        request.BankData,
	}
	if err := db.Create(&employee).Error; err != nil {
		logger.Errorf("Error to create employee: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating employee on database")
		return
	}

	sendSuccess(ctx, "employee created", employee)
}
