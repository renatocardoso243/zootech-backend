package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateAnimalRequest
type CreateAnimalRequest struct {
	Name        string  `json:"name"`                       // Nome do animal
	Tag         string  `json:"tag"`                        // Identificador do animal
	Race        string  `json:"race"`                       // Raça do animal
	Weight      int64   `json:"weight"`                     // Peso do animal
	Sex         string  `json:"sex"`                        // Sexo do animal
	Birthdate   string  `json:"birth_date"`                 // Data de nascimento do animal
	Vaccinated  string   `json:"vaccinated"`                 // Indica se o animal foi vacinado
	Health    	string  `json:"health"`                     // Status de saúde do animaluint    `json:"health_id"`                  // ID para a tabela de status de saúde
	HerdID      uint    `json:"herd_id"`                    // ID para a tabela de rebanho
	Observation string  `json:"observation"`                // Observações gerais

}

// Validates CreateAnimalRequest
func (r *CreateAnimalRequest) Validate() error {
	if r.Name == "" && r.Tag == "" && r.Race == "" && r.Weight == 0 && r.Sex == "" && r.Birthdate == "" && r.Vaccinated == "" && r.Health == "" && r.HerdID == 0 && r.Observation == ""  {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Tag == "" {
		return errParamIsRequired("tag", "string")
	}
	if r.Race == "" {
		return errParamIsRequired("race", "string")
	}
	if r.Weight == 0 {
		return errParamIsRequired("weight", "int64")
	}
	if r.Sex == "" {
		return errParamIsRequired("sex", "string")
	}
	if r.Birthdate == "" {
		return errParamIsRequired("birth_date", "string")
	}
	if r.Vaccinated == "" {
		return errParamIsRequired("vaccinated", "string")
	}
	if r.Health == "" {
		return errParamIsRequired("health_id", "uint")
	}
	if r.HerdID == 0 {
		return errParamIsRequired("herd_id", "uint")
	}
	if r.Observation == "" {
		return errParamIsRequired("observation", "string")
	}
	return nil
}

// UpdateAnimalRequest
type UpdateAnimalRequest struct {
	Name        string  `json:"name"`                  
	Tag         string  `json:"tag"`
	Race        string  `json:"race"`                  
	Weight      int64   `json:"weight"`                
	Sex         string  `json:"sex"`                   
	Birthdate   string  `json:"birth_date"`                 
	Vaccinated  string   `json:"vaccinated"`                 
	Health    	string  `json:"health"`
	HerdID      uint    `json:"herd_id"`                    
	Observation string  `json:"observation"`               
}

// Validates UpdateAnimalRequest
func (r *UpdateAnimalRequest) Validate() error {
	// If any field provided, validation is true
	if r.Name != "" || r.Tag != "" || r.Race != "" || r.Weight != 0 || r.Sex != "" || r.Birthdate != "" || r.Vaccinated != "" || r.Health != "" || r.HerdID != 0 || r.Observation != ""  {
		return nil
	}
	// If all fields are empty. return false
	return fmt.Errorf("request body is empty or malformed")
}

// CreateHerdRequest
type CreateHerdRequest struct {
	HerdName    string  `json:"herd_name"`                  // Nome do rebanho
	Type        string  `json:"type"`                       // Tipo de rebanho
	Description string  `json:"description"`                // Descrição do rebanho
}

// Create herd request validation
func (r *CreateHerdRequest) Validate() error {
	if r.HerdName == "" && r.Type == "" && r.Description == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.HerdName == "" {
		return errParamIsRequired("herd_name", "string")
	}
	if r.Type == "" {
		return errParamIsRequired("type", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}

// UpdateHerdRequest
type UpdateHerdRequest struct {
	HerdName    string  `json:"herd_name"`                  // Nome do rebanho
	Type        string  `json:"type"`                       // Tipo de rebanho
	Description string  `json:"description"`                // Descrição do rebanho
}

// Validates UpdateHerdRequest
func (r *UpdateHerdRequest) Validate() error {
	// If any field provided, validation is true
	if r.HerdName != "" || r.Type != "" || r.Description != ""  {
		return nil
	}
	// If all field or a single field was empty. return false
	return fmt.Errorf("request body is empty or malformed")
}

// Create individual diet request
type CreateDietRequest struct {
	DietName    	string    `json:"diet_name"`
	AnimalType  	string    `json:"animal_type"`
	Objective    	string    `json:"objective"`
	Ingredients 	string    `json:"ingredients"`
	NutritionalInfo string 	  `json:"nutritional_info"`
}

// Create individual diet request validation
func (r *CreateDietRequest) Validate() error {
	if r.DietName == "" && r.AnimalType == "" && r.Objective == "" && r.Ingredients == "" && r.NutritionalInfo == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.DietName == "" {
		return errParamIsRequired("diet_name", "string")
	}
	if r.AnimalType == "" {
		return errParamIsRequired("type", "string")
	}
	if r.Objective == "" {
		return errParamIsRequired("start_date", "string")
	}
	if r.Ingredients == "" {
		return errParamIsRequired("end_date", "string")
	}
	if r.NutritionalInfo == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}


// Update individual diet request
type UpdateDietRequest struct {
	DietName    	string    `json:"diet_name"`
	AnimalType  	string    `json:"animal_type"`
	Objective    	string    `json:"objective"`
	Ingredients 	string    `json:"ingredients"`
	NutritionalInfo string 	  `json:"nutritional_info"`
}

// Update individual diet request validation
func (r *UpdateDietRequest) Validate() error {
	// If any field provided, validation is true
	if r.DietName != "" || r.AnimalType != "" || r.Objective != "" || r.Ingredients != "" || r.NutritionalInfo != "" {
		return nil
	}
	// If all field or a single field was empty. return false
	return fmt.Errorf("request body is empty or malformed")
}



// Create weight request
type CreateWeightRequest struct {
	Data 		string 		`json:"data"`
	TotalWeight string 		`json:"total_weight"`
	Gain 		string 		`json:"gain"`
	Loss 		string 		`json:"loss"`
	AnimalID 	uint 		`json:"animal_id"`
}

// Create weight request validation
func (r *CreateWeightRequest) Validate() error {
	if r.Data == "" && r.TotalWeight == "" && r.Gain == "" && r.Loss == "" && r.AnimalID == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Data == "" {
		return errParamIsRequired("data", "string")
	}
	if r.TotalWeight == "" {
		return errParamIsRequired("total_weight", "int64")
	}
	if r.Gain == "" {
		return errParamIsRequired("gain", "int64")
	}
	if r.Loss == "" {
		return errParamIsRequired("loss", "int64")
	}
	if r.AnimalID == 0 {
		return errParamIsRequired("animal_id", "uint")	
	}
	return nil
}

// Update weight request
type UpdateWeightRequest struct {
	Data 		string 		`json:"data"`
	TotalWeight string 		`json:"total_weight"`
	Gain 		string 		`json:"gain"`
	Loss 		string		`json:"loss"`
	AnimalID 	uint 		`json:"animal_id"`
}

// Update weight request validation
func (r *UpdateWeightRequest) Validate() error {
	// If any field provided, validation is true
	if r.Data != "" || r.TotalWeight != "" || r.Gain != "" || r.Loss != "" || r.AnimalID != 0 {
		return nil
	}
	// If all field or a single field was empty. return false
	return fmt.Errorf("request body is empty or malformed")
}
