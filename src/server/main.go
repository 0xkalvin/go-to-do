package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-to-do/src/database"
	"go-to-do/src/controllers/todos"
)

func SetupRouter(database *gorm.DB) *gin.Engine {
	router := gin.Default()
	
	router.Use(func(c *gin.Context) {
		c.Set("database", database)
		c.Next()
	})
	
	router.POST("/todos", todos.Create)
	router.GET("/todos", todos.Index)
	router.GET("/todos/:id", todos.Show)
	router.PUT("/todos/:id", todos.Update)
	router.DELETE("/todos/:id", todos.Delete)

	return router
}

func main() {

	var database = database.InitializeDatabase()
	var router = SetupRouter(database)
	
	router.Run(":3000") 	
}