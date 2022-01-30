package dto

type CreateEmployeeInput struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Department string `json:"department" binding:"required"`
}

type UpdateEmployeeInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Department string `json:"department"`
}
