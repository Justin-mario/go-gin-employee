package validator

type CreateEmployeeInput struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Department string `json:"department" binding:"required"`
}
