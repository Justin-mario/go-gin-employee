package main

import (
	"github.com/Justin-mario/go-gin-employee/entity"
	"github.com/Justin-mario/go-gin-employee/handler"
	"github.com/Justin-mario/go-gin-employee/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	middlewares.SetUpLogOutPut()
	entity.ConnectDatabase()

	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.POST("/", handler.CreateEmployee)
		apiRoutes.PATCH("/:id", handler.UpdateEmployee)
		apiRoutes.DELETE("/:id", handler.DeleteEmployee)
		apiRoutes.DELETE("/", handler.DeleteAllEmployee)
	}

	apiView := server.Group("/")
	{
		apiView.GET("/", handler.FindEmployees)
		apiView.GET("/:id", handler.FindEmployee)
	}

	err := server.Run()
	if err != nil {
		return
	}
}
