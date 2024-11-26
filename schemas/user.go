package schemas

// Estrutura do usuário
type User struct {
	ID        uint   `gorm:"primaryKey"` // Chave primária
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique;not null"` // Campo único e obrigatório
	Password  string `json:"password"`
}