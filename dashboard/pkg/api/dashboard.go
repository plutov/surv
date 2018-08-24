package api

// DashboardRow definition contains single survey submission
type DashboardRow struct {
	SurveyServiceName string            `json:"survey_service_name"`
	SurveyID          int64             `json:"survey_id"`
	AnswerID          string            `json:"answer_id"`
	Values            map[string]string `json:"values"`
	User              string            `json:"user"`
}

// SurveyService definition
type SurveyService struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
