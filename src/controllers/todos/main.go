package todos

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-to-do/src/models"	
)

func Create(c *gin.Context) {

	var todo models.Todo
	c.BindJSON(&todo)

	database := c.MustGet("database").(*gorm.DB)
	err := database.Create(&todo).Error	

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} 
	
	c.JSON(http.StatusCreated, todo)
}

func Delete(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	database := c.MustGet("database").(*gorm.DB)
	err := database.Where("id = ?", id).Delete(&todo).Error; 

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo " + id + " has been deleted",
	})
}

func Index(c *gin.Context) {

	var todos []models.Todo
	
	database := c.MustGet("database").(*gorm.DB)
	err := database.Find(&todos).Error

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func Show(c *gin.Context) {

	var todo models.Todo
	id := c.Param("id")
	
	database := c.MustGet("database").(*gorm.DB)
	err := database.First(&todo, id).Error; 

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func Update(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	
	database := c.MustGet("database").(*gorm.DB)
	err := database.First(&todo, id).Error; 

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.BindJSON(&todo)
	err = database.Save(&todo).Error

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todo)
}