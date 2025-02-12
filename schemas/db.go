package schemas

import (
	"gorm.io/gorm"
)

// Estrutura do usuário
type User struct {
	ID        uint   `gorm:"primaryKey"` // Chave primária
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique;not null"` // Campo único e obrigatório
	Password  string `json:"password"`
}

// Estrutura dos funcionários
type Employee struct {
	ID              uint   `gorm:"primaryKey"`  // Chave primária
	EmployeeId      string `json:"employee_id"` // ID do funcionário
	FullName        string `json:"full_name"`
	Birthdate       string `json:"birth_date"`
	Genre           string `json:"genre"`
	CivilStatus     string `json:"civil_status"`
	Nacionality     string `json:"nacionality"`
	TaxIdNumber     string `json:"tax_id_number" gorm:"unique;not null"` // CPF Campo único e obrigatório
	IdentityCard    string `json:"identity_card"`                        //RG
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	AdmissionDate   string `json:"admission_date"`
	Departament     string `json:"departament"`
	ResignationDate string `json:"resignation_date"`
	WorkRegiment    string `json:"work_regiment"`
	Salary          string `json:"salary"`
	BankData        string `json:"bank_data"`
	Task            []Task `json:"task,omitempty" gorm:"foreignKey:EmployeeID"` // Relacionamento com atividades

}

// Estrutura para as atividades
type Task struct {
	ID             uint   `gorm:"primaryKey"`  // Chave primária
	EmployeeID     string `json:"employee_id"` // Chave estrangeira para Employee
	TaskName       string `json:"task_name"`
	TaskDate       string `json:"task_date"`
	TaskTime       string `json:"task_time"`
	TaskConclusion bool   `json:"task_conclusion"`

	Employee *Employee `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"` // Relacionamento com Employee
}

// Estrutura para os animais
type Animal struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Tag         string `json:"tag" gorm:"unique"`
	Race        string `json:"race"`
	Weight      int64  `json:"weight"`
	Sex         string `json:"sex"`
	Birthdate   string `json:"birth_date"`
	Vaccinated  string `json:"vaccinated"`
	Health      string `json:"health"`
	HerdID      uint   `json:"herd_id"` // Chave estrangeira para Herd
	Observation string `json:"observation"`

	Herd *Herd `json:"herd,omitempty" gorm:"foreignKey:HerdID"` // Relacionamento com Herd
}

// Estrutura para rebanhos
type Herd struct {
	gorm.Model
	ID          uint     `gorm:"primaryKey"`
	HerdName    string   `json:"herd_name"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Animals     []Animal `json:"animals" gorm:"foreignKey:HerdID"` // Relacionamento reverso (opcional)
}

// Dieta individual para um animal
type Diet struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey"`
	DietName        string `json:"diet_name"`
	AnimalType      string `json:"animal_type"`
	Objective       string `json:"objective"`
	Ingredients     string `json:"ingredients"`
	NutritionalInfo string `json:"nutritional_info"`
}

type Weight struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Data        string `json:"data"`
	TotalWeight string `json:"total_weight"`
	Gain        string `json:"gain"`
	Loss        string `json:"loss"`
	AnimalID    uint   `json:"animal_id"`

	Animal *Animal `json:"animal,omitempty" gorm:"foreignKey:AnimalID"`
}
