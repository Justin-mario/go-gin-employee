package controllers

import (
	"github.com/Justin-mario/go-gin-employee/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateEmployeeInput struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Department string `json:"department" binding:"required"`
}

func FindEmployees(c *gin.Context) {
	var employees []entity.Employee
	entity.DB.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func CreateEmployee(c *gin.Context) {
	var input CreateEmployeeInput

	//if err := c.ShouldBindJSON(&input)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error()})
	//	return
	//}

	employee := entity.Employee{FirstName: input.FirstName, LastName: input.LastName, Department: input.Department, DateOfEmployment: time.Now()}
	entity.DB.Create(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}
