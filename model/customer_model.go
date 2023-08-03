package model

import "time"

type CustomerModel struct {
	Id        int
	User_id   int
	Nik       string
	Name      string
	Email     string
	Phone     string
	Address   string
	Birthdate time.Time
	Status    string
}
