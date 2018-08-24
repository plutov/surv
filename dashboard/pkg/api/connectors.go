package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Connector interface
type Connector interface {
	GetAnswers(string, string) ([]DashboardRow, error)
}

// GetConnectorInstanceByName func
func GetConnectorInstanceByName(name string) (Connector, error) {
	switch name {
	case "surv":
		return new(SurvConnector), nil
	}

	return nil, errors.New("connector is not implemented")
}

// SurvConnector implementation
type SurvConnector struct{}

// SurvAnswer definition
type SurvAnswer struct {
	ID       string            `json:"id"`
	SurveyID int64             `json:"survey_id"`
	Values   map[string]string `json:"values"`
	User     string            `json:"user"`
}

// GetAnswers from SURV services
func (c *SurvConnector) GetAnswers(name string, address string) ([]DashboardRow, error) {
	rows := []DashboardRow{}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	r, err := httpClient.Get(address + "/answers")
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
			SurveyServiceName: name,
			SurveyID:          a.SurveyID,
			AnswerID:          a.ID,
			Values:            a.Values,
			User:              a.User,
		})
	}

	return rows, nil
}
