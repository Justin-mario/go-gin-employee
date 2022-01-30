package main

import (
	"github.com/Justin-mario/go-gin-employee/controllers"
	"github.com/Justin-mario/go-gin-employee/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	entity.ConnectDatabase()

	server.GET("/", controllers.FindEmployees)
	server.POST("/", controllers.CreateEmployee)

	server.Run()
}
