package api

import (
	"encoding/json"
	"net/http"
	"time"
)

// Connector interface
type Connector interface {
	GetAnswers() ([]DashboardRow, error)
}

// SurvConnector implementation
type SurvConnector struct {
	Name    string
	Address string
}

// SurvAnswer definition
type SurvAnswer struct {
	ID       string            `json:"id"`
	SurveyID int64             `json:"survey_id"`
	Values   map[string]string `json:"values"`
	User     string            `json:"user"`
}

// GetAnswers from SURV services
func (c *SurvConnector) GetAnswers() ([]DashboardRow, error) {
	rows := []DashboardRow{}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	r, err := httpClient.Get(c.Address + "/answers")
	if err != nil {
		return rows, err
	}

	defer r.Body.Close()

	var answers []SurvAnswer
	err = json.NewDecoder(r.Body).Decode(&answers)
	if err != nil {
		return rows, err
	}

	for _, a := range answers {
		rows = append(rows, DashboardRow{
			SurveyServiceName: c.Name,
			SurveyID:          a.SurveyID,
			AnswerID:          a.ID,
			Values:            a.Values,
			User:              a.User,
		})
	}

	return rows, nil
}
