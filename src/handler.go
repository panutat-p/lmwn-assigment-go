package src

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"API version": "1.0.0",
	})
}

func Persons(c *gin.Context) {
	persons, err := GetPersons()
	if err != nil {
		log.Println("Failed to get COVID19 report, err:", err)
		c.JSON(500, gin.H{
			"message": "Failed to get COVID19 report",
		})
		return
	}

	c.JSON(200, gin.H{
		"persons": persons,
	})
}
