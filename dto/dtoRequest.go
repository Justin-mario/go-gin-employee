package dto

type CreateEmployeeInput struct {
	FirstName  string `json:"first_name" binding:"required,min=2,max=20"`
	LastName   string `json:"last_name" binding:"required,min=2,max=20"`
	Age        int8   `json:"age" binding:"required,gte=18,lte=70"`
	Email      string `json:"email" binding:"required,email"`
	Department string `json:"department" binding:"required"`
}

type UpdateEmployeeInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Department string `json:"department"`
}
