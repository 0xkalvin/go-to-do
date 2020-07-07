package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-to-do/src/database"
	"go-to-do/src/controllers/todos"
	"go-to-do/src/middlewares"
)

func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Up and kicking"})
}

func SetupRouter(database *gorm.DB) *gin.Engine {
	router := gin.Default()
	
	router.Use(func(c *gin.Context) {
		c.Set("database", database)
		c.Next()
	})

	router.Use(middlewares.AddRequestId())

	router.GET("/health", HealthHandler)

	v1 := router.Group("/v1")
	{
		v1.POST("/todos", todos.Create)
		v1.GET("/todos", todos.Index)
		v1.GET("/todos/:id", todos.Show)
		v1.PUT("/todos/:id", todos.Update)
		v1.DELETE("/todos/:id", todos.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		message := "Resource `" + path + "` not found"
		c.JSON(404, gin.H{"message": message})
	})

	return router
}

func main() {

	var database = database.InitializeDatabase()
	var router = SetupRouter(database)
	
	log.Println("Starting server on port 3000")
	router.Run(":3000") 	
}