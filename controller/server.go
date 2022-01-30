package main

import (
	"github.com/Justin-mario/go-gin-employee/entity"
	"github.com/Justin-mario/go-gin-employee/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	entity.ConnectDatabase()

	server.GET("/", handler.FindEmployees)
	server.GET("/:id", handler.FindEmployee)
	server.POST("/", handler.CreateEmployee)
	server.PATCH("/:id", handler.UpdateEmployee)
	server.DELETE("/:id", handler.DeleteEmployee)
	server.DELETE("/", handler.DeleteAllEmployee)

	err := server.Run()
	if err != nil {
		return
	}
}
