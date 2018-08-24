package api

// Answer definition
type Answer struct {
	ID       string            `json:"id"`
	SurveyID int64             `json:"survey_id"`
	Values   map[string]string `json:"values"`
	User     string            `json:"user"`
}
