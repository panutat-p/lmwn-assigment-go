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
	persons, err := GetCovidReport()
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

func GroupByAge(c *gin.Context) {
	persons, err := GetCovidReport()
	if err != nil {
		log.Println("Failed to get COVID19 report, err:", err)
		c.JSON(500, gin.H{
			"message": "Failed to get COVID19 report",
		})
		return
	}

	var (
		summary = make(map[string]int)
	)

	for _, e := range persons {
		if e.Province == "Nan" {
			e.Province = ""
		}
		if e.ProvinceEn == "Nan" {
			e.ProvinceEn = ""
		}

		if e.Province != "" {
			summary[e.Province] += e.Age
		}
	}

	c.JSON(200, gin.H{
		"length":  len(summary),
		"summary": summary,
	})
}
