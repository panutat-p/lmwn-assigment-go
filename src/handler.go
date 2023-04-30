package src

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	report Reporter
}

func NewHandler(r Reporter) *Handler {
	return &Handler{
		report: r,
	}
}

func (h Handler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"API version": "1.0.0",
	})
}

func (h Handler) GetReport(c *gin.Context) {
	report, err := h.report.GetCovidReport()
	if err != nil {
		log.Println("Failed to get COVID19 report, err:", err)
		c.JSON(500, gin.H{
			"message": "Failed to get COVID19 report",
		})
		return
	}

	c.JSON(200, report)
}

func (h Handler) GetSummary(c *gin.Context) {
	report, err := h.report.GetCovidReport()
	if err != nil {
		log.Println("Failed to get COVID19 report, err:", err)
		c.JSON(500, gin.H{
			"message": "Failed to get COVID19 report",
		})
		return
	}

	param := []string{
		"Nakhon Ratchasima",
		"Nong Khai",
	}
	summary := GenerateSummary(report, NewSetFromSlice(param))

	c.JSON(200, summary)
}
