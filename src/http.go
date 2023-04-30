package src

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

//go:generate mockgen -destination=../src/mock_reporter.go -package=src lmwn-assignment-go/src Reporter
type Reporter interface {
	GetCovidReport() ([]*Person, error)
}

type Report struct {
	c       http.Client
	baseURL string
}

func NewReport(url string, timeout time.Duration) *Report {
	return &Report{
		c:       http.Client{Timeout: timeout},
		baseURL: url,
	}
}

func (r Report) GetCovidReport() ([]*Person, error) {
	var (
		report Response
	)

	resp, err := r.c.Get(r.baseURL)
	if err != nil {
		return nil, err
	}
	defer CloseBody(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&report)
	if err != nil {
		return nil, err
	}

	return report.Data, nil
}

func CloseBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Println("Failed to close ReadCloser, err:", err)
	}
}
