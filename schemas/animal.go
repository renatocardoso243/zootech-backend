package schemas

import (
	"gorm.io/gorm"
)

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
type IndividualDiet struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"` 
	DietName    string    `json:"diet_name"`
	Type        string    `json:"type"`
	StartDate   string 	  `json:"start_date"`
	EndDate     string 	  `json:"end_date"`
	Description string    `json:"description"`
	AnimalID    uint      `json:"animal_id"`  // Chave estrangeira para Animal

	Animal      *Animal   `json:"animal,omitempty" gorm:"foreignKey:AnimalID"` // Relacionamento com Animal
}

// Dieta de grupo para um rebanho
type GroupDiet struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"` 
	DietName    string    `json:"diet_name"`
	Type        string    `json:"type"`
	StartDate   string 	  `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Description string    `json:"description"`
	HerdID      uint      `json:"herd_id"`    // Chave estrangeira para Herd
	
	Herd        *Herd     `json:"herd,omitempty" gorm:"foreignKey:HerdID"` // Relacionamento com Herd
}
