package main

import (
	"github.com/Justin-mario/go-gin-employee/entity"
	"github.com/Justin-mario/go-gin-employee/handler"
	"github.com/Justin-mario/go-gin-employee/middlewares"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	middlewares.SetUpLogOutPut()
	entity.ConnectDatabase()

	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.POST("/add", handler.CreateEmployee)
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
