package models

import "time"

type User struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Contract struct {
	ID               string    `json:"id"`
	PlanID           string    `json:"planid"`
	UserID           string    `json:"userid"`
	DateOfContract   time.Time `json:"datecon"`
	DateOfExpiration time.Time `json:"dateexp"`
}

type NewUser struct {
	ID        string
	Nombre    string
	Apellidos string
	UserName  string
	Email     string
	Password  string
}
