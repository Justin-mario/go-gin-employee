package entity

import "time"

type Employee struct {
	ID               uint      `json:"id" gorm:"primary_key"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Department       string    `json:"department"`
	DateOfEmployment time.Time `json:"dateOfEmployment"`
}
