package api

import "testing"

func TestSave(t *testing.T) {
	s := NewStorage()
	s.Save(Answer{
		SurveyID: 1,
	})
	s.Save(Answer{
		SurveyID: 1,
	})

	if len(s.data) != 2 {
		t.Fatalf("expected 2 items in storage, got %d", len(s.data))
	}
}
