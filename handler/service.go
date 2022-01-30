package handler

import (
	"fmt"
	"github.com/Justin-mario/go-gin-employee/dto"
	"github.com/Justin-mario/go-gin-employee/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func FindEmployees(c *gin.Context) {
	var employees []entity.Employee
	entity.DB.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func CreateEmployee(c *gin.Context) {
	var input dto.CreateEmployeeInput
	fmt.Println("user input", input)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	employee := entity.Employee{FirstName: input.FirstName, LastName: input.LastName, Department: input.Department, DateOfEmployment: time.Now()}
	fmt.Println("user input", input)

	fmt.Println("employee object", employee)
	entity.DB.Create(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func FindEmployee(c *gin.Context) {
	var employee entity.Employee
	if err := entity.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func UpdateEmployee(c *gin.Context) {
	var employee entity.Employee
	if err := entity.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record Not Found!"})
		return
	}

	var input dto.UpdateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	entity.DB.Model(&employee).Update(input)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func DeleteEmployee(c *gin.Context) {
	var employee entity.Employee
	if err := entity.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record Not Found!"})
		return
	}

	entity.DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"Deleted": true})
}

func DeleteAllEmployee(c *gin.Context) {
	var employees []entity.Employee
	entity.DB.Find(&employees)
	if len(employees) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Employee DataBase Is Empty!"})
		return
	}
	entity.DB.Delete(&employees)
	c.JSON(http.StatusOK, gin.H{"Deleted": true})
}
