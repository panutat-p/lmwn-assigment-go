package src

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	Client = http.Client{
		Timeout: 5 * time.Second,
	}
)

const (
	ReportURL = "https://static.wongnai.com/devinterview/covid-cases.json"
)

func GetPersons() ([]Person, error) {
	var (
		report Response
	)

	resp, err := Client.Get(ReportURL)
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
