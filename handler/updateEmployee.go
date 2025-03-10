package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatocardoso243/gopportunities.git/schemas"
)

func UpdateEmployeeHandler(ctx *gin.Context) {
	request := UpdateEmployeeRequest{}

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

	employee := schemas.Employee{}
	if err := db.First(&employee, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "employee not found")
		return
	}

	// Update employee
	if request.EmployeeId != 0 {
		employee.EmployeeId = request.EmployeeId
	}
	if request.FullName != "" {
		employee.FullName = request.FullName
	}
	if request.Birthdate != "" {
		employee.Birthdate = request.Birthdate
	}
	if request.Genre != "" {
		employee.Genre = request.Genre
	}
	if request.CivilStatus != "" {
		employee.CivilStatus = request.CivilStatus
	}
	if request.Nacionality != "" {
		employee.Nacionality = request.Nacionality
	}
	if request.TaxIdNumber != "" {
		employee.TaxIdNumber = request.TaxIdNumber
	}
	if request.IdentityCard != "" {
		employee.IdentityCard = request.IdentityCard
	}
	if request.Address != "" {
		employee.Address = request.Address
	}
	if request.Phone != "" {
		employee.Phone = request.Phone
	}
	if request.Email != "" {
		employee.Email = request.Email
	}
	if request.Role != "" {
		employee.Role = request.Role
	}
	if request.AdmissionDate != "" {
		employee.AdmissionDate = request.AdmissionDate
	}
	if request.Departament != "" {
		employee.Departament = request.Departament
	}
	if request.WorkRegiment != "" {
		employee.WorkRegiment = request.WorkRegiment
	}
	if request.ResignationDate != "" {
		employee.ResignationDate = request.ResignationDate
	}
	if request.Salary != "" {
		employee.Salary = request.Salary
	}
	if request.BankData != "" {
		employee.BankData = request.BankData
	}

	if err := db.Save(&employee).Error; err != nil {
		logger.Errorf("Error to update employee: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating employee on database")
		return
	}

	sendSuccess(ctx, "employee updated", employee)
}
