package day2

import "testing"

func TestSafeReports(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 2,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 490,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reports, err := parseReports(tt.data)
			if err != nil {
				t.Fatalf("failed to parse games %v", err)
			}
			value := reports.calculateSafeReports()
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}

func TestSafeDampenedReports(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 4,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 536,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reports, err := parseReports(tt.data)
			if err != nil {
				t.Fatalf("failed to parse games %v", err)
			}
			value := reports.calculateSafeDampenedReports()
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}
