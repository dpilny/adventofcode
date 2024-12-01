package day2

import "testing"

func TestGame(t *testing.T) {
	tests := []struct {
		name  string
		data  string
		red   int
		green int
		blue  int
		want  int
	}{
		{
			name:  "AoC Example",
			data:  "sample",
			red:   12,
			green: 13,
			blue:  14,
			want:  8,
		},
		{
			name:  "AoC Task",
			data:  "task",
			red:   12,
			green: 13,
			blue:  14,
			want:  3035,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := sumValidGames(tt.data, tt.red, tt.green, tt.blue)
			if err != nil {
				t.Fatalf("failed to parse games %v", err)
			}
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}
