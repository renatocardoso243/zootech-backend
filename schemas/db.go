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

// Estrutura para os animais
type Animal struct {
	gorm.Model
	ID			uint		   `gorm:"primaryKey"`                 
	Name        string         `json:"name"`
	Tag         string         `json:"tag" gorm:"unique"`
	Race        string         `json:"race"`
	Weight      int64          `json:"weight"`
	Sex         string         `json:"sex"`
	Birthdate   string         `json:"birth_date"`
	Vaccinated  string         `json:"vaccinated"`       
	Health    	string         `json:"health"`         
	HerdID      uint           `json:"herd_id"`           // Chave estrangeira para Herd
	Observation string         `json:"observation"`
	
	Herd   		*Herd         `json:"herd,omitempty" gorm:"foreignKey:HerdID"`     // Relacionamento com Herd
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
	ID          	uint      `gorm:"primaryKey"` 
	DietName    	string    `json:"diet_name"`
	AnimalType  	string    `json:"animal_type"`
	Objective    	string    `json:"objective"`
	Ingredients 	string    `json:"ingredients"`
	NutritionalInfo string 	  `json:"nutritional_info"`
}

type Weight struct {
	gorm.Model
	ID 			uint 		`gorm:"primaryKey"`
	Data 		string 		`json:"data"`
	TotalWeight string 		`json:"total_weight"`
	Gain 		string 		`json:"gain"`
	Loss 		string		`json:"loss"`
	AnimalID    uint		`json:"animal_id"`

	Animal 		*Animal 	`json:"animal,omitempty" gorm:"foreignKey:AnimalID"`

}

