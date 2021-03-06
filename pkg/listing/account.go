package listing

import "time"

type Account struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CPF       string     `json:"cpf,omitempty"`
	Secret    string     `json:"-"`
	Balance   float64    `json:"balance"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
