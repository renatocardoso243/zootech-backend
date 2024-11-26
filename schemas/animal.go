package schemas

import (
	"time"

	"gorm.io/gorm"
)

// Estrutura para os animais
type Animal struct {
	gorm.Model                 // Inclui campos como ID, CreatedAt, UpdatedAt e DeletedAt
	Name        string         `json:"name"`
	Tag         string         `json:"tag" gorm:"unique"` // Tag deve ser única
	Race        string         `json:"race"`
	Weight      int64          `json:"weight"`
	Sex         string         `json:"sex"`
	Birthdate   string      	`json:"birthdate"`
	Vaccinated  string          `json:"vaccinated"`        // Nome corrigido para "Vaccinated"
	Health    	string          `json:"health"`         // Chave estrangeira para HealthStatus
	HerdID      uint           `json:"herd_id"`           // Chave estrangeira para Herd
	Observation string         `json:"observation"`

	// Relacionamentos 
	Herd   *Herd         `json:"herd,omitempty" gorm:"foreignKey:HerdID"`     // Relacionamento com Herd
}


// Estrutura para rebanhos
type Herd struct {
	ID          uint    `gorm:"primaryKey"` // Chave primária
	HerdName    string  `json:"herd_name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Animals     []Animal `json:"animals" gorm:"foreignKey:HerdID"` // Relacionamento reverso (opcional)
}

// Dieta individual para um animal
type IndividualDiet struct {
	ID          uint      `gorm:"primaryKey"` // Chave primária
	DietName    string    `json:"diet_name"`
	Type        string    `json:"type"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	AnimalID    uint      `json:"animal_id"`  // Chave estrangeira para Animal
	Animal      *Animal   `json:"animal,omitempty" gorm:"foreignKey:AnimalID"` // Relacionamento com Animal
}

// Dieta de grupo para um rebanho
type GroupDiet struct {
	ID          uint      `gorm:"primaryKey"` // Chave primária
	DietName    string    `json:"diet_name"`
	Type        string    `json:"type"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	HerdID      uint      `json:"herd_id"`    // Chave estrangeira para Herd
	Herd        *Herd     `json:"herd,omitempty" gorm:"foreignKey:HerdID"` // Relacionamento com Herd
}
