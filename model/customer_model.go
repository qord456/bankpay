package model

import "time"

type CustomerModel struct {
	Id        int
	UserId    int
	Nik       string
	Name      string
	Email     string
	Phone     string
	Address   string
	Birthdate time.Time
	Balance   int
	Status    string
}
