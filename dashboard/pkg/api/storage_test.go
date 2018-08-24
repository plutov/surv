package api

import "testing"

func TestSave(t *testing.T) {
	s := NewStorage()
	s.Save(DashboardRow{
		AnswerID: "1",
	})
	s.Save(DashboardRow{
		AnswerID: "1",
	})

	// check that we don't save same answer ID 2 times
	if len(s.data) != 1 {
		t.Fatalf("expected 1 item in storage, got %d", len(s.data))
	}
}

func TestGet(t *testing.T) {
	s := NewStorage()
	s.Save(DashboardRow{
		AnswerID: "1",
	})
	s.Save(DashboardRow{
		AnswerID: "2",
	})

	var tests = []struct {
		name    string
		limit   int
		offset  int
		wantLen int
	}{
		{"0, 0", 0, 0, 0},
		{"0, 1", 0, 1, 0},
		{"1, 0", 1, 0, 1},
		{"2, 2", 2, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Get(tt.limit, tt.offset)
			if len(got) != tt.wantLen {
				t.Fatalf("Get(%d, %d) got %d items, want %d", tt.limit, tt.offset, len(got), tt.wantLen)
			}
		})
	}
}
