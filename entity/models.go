package entity

import "time"

type Employee struct {
	ID               uint      `json:"id" gorm:"primary_key"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Age              int8      `json:"age"`
	Email            string    `json:"email"`
	Department       string    `json:"department"`
	DateOfEmployment time.Time `json:"dateOfEmployment"`
}
