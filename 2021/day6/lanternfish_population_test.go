package day6

import "testing"

func TestPopulationSimulation80(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 5934,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 374927,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			population, err := parseInitialPopulation(tt.data)
			if err != nil {
				t.Fatalf("failed to parse population %v", err)
			}
			resultValue := population.countPopulationAfterDays(80)
			if err != nil {
				t.Fatalf("failed get population %v", err)
			}
			if resultValue != tt.want {
				t.Errorf("population = %v, want %v", resultValue, tt.want)
			}
		})
	}
}

func TestPopulationSimulation256(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 26984457539,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 374927,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			population, err := parseInitialPopulation(tt.data)
			if err != nil {
				t.Fatalf("failed to parse population %v", err)
			}
			resultValue := population.countPopulationAfterDays(256)
			if err != nil {
				t.Fatalf("failed get population %v", err)
			}
			if resultValue != tt.want {
				t.Errorf("population = %v, want %v", resultValue, tt.want)
			}
		})
	}
}
