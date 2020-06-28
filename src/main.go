package main

import (
	"github.com/gin-gonic/gin"
	"go-to-do/src/database"
	"go-to-do/src/controllers/todos"
)

func main() {

	router := gin.Default()
	var database = database.InitializeDatabase()
	
	router.Use(func(c *gin.Context) {
		c.Set("database", database)
		c.Next()
	})
	
	router.POST("/todos", todos.Create)
	router.GET("/todos", todos.Index)
	router.GET("/todos/:id", todos.Show)
	router.PUT("/todos/:id", todos.Update)
	router.DELETE("/todos/:id", todos.Delete)
	
	router.Run(":3000") 	
}