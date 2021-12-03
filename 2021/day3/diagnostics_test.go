package day3

import "testing"

func TestPowerConsumption(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 198,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 693486,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagData, err := parseDiagnosticInput(tt.data)
			if err != nil {
				t.Fatalf("failed to parse diagData")
			}
			consumption, err := calculatePowerConsumption(diagData)
			if err != nil {
				t.Fatalf("failed to calculate power consumption %v", err)
			}
			if consumption != tt.want {
				t.Errorf("Consumption = %v, want %v", consumption, tt.want)
			}
		})
	}
}

func TestLifeSupportRating(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 230,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 3379326,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagData, err := parseDiagnosticInput(tt.data)
			if err != nil {
				t.Fatalf("failed to parse diagData")
			}
			lifeSupportRating, err := calculateLifeSupportRating(diagData)
			if err != nil {
				t.Fatalf("failed to calculate power lifeSupportRating %v", err)
			}
			if lifeSupportRating != tt.want {
				t.Errorf("Consumption = %v, want %v", lifeSupportRating, tt.want)
			}
		})
	}
}
