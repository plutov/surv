package api

import "testing"

func TestGetConnectorInstanceByName(t *testing.T) {
	_, err := GetConnectorInstanceByName("surv")
	if err != nil {
		t.Fatalf("not expected err, got %v", err)
	}

	_, err = GetConnectorInstanceByName("typeform")
	if err == nil {
		t.Fatal("expected err, got nil")
	}
}
