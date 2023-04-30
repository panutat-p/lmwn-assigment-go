package main

import (
	"github.com/gin-gonic/gin"
	"lmwn-assignment-go/src"
	"log"
	"time"
)

const (
	Port = "8080"
)

func main() {
	report := src.NewReport("https://static.wongnai.com/devinterview/covid-cases.json", 5*time.Second)
	handler := src.NewHandler(report)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", handler.Health)
	router.GET("/report", handler.GetReport)
	router.GET("/summary", handler.GetSummary)
	log.Println("Web server is listening on port", Port)
	err := router.Run(":" + Port)
	if err != nil {
		log.Println("Failed to run the web server, err:", err)
	}
}
