package models

// Password represent format in update-password proccess
type Password struct {
	NewPasword     string `json:"new_password"`
	ActualPassword string `json:"actual_password"`
}
