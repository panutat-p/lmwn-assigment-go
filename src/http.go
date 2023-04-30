package src

import (
	"encoding/json"
	"net/http"
	"time"
)

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
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&report)
	if err != nil {
		return nil, err
	}

	return report.Data, nil
}
