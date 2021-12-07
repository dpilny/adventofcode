package day7

import "testing"

func TestPopulationSimulationMean(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 37,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 364898,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crabPositions, err := parseCrabAlignment(tt.data)
			if err != nil {
				t.Fatalf("failed to parse crabPositions %v", err)
			}
			alignmentFuel := calculateAlignLeastFuel(crabPositions)
			if err != nil {
				t.Fatalf("failed calculate alignmentFuel %v", err)
			}
			if alignmentFuel != tt.want {
				t.Errorf("alignmentFuel = %v, want %v", alignmentFuel, tt.want)
			}
		})
	}
}

func TestPopulationSimulationAvg(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 168,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 104149091,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crabPositions, err := parseCrabAlignment(tt.data)
			if err != nil {
				t.Fatalf("failed to parse crabPositions %v", err)
			}
			alignmentFuel := calculateAlignBruteForce(crabPositions)
			if err != nil {
				t.Fatalf("failed calculate alignmentFuel %v", err)
			}
			if alignmentFuel != tt.want {
				t.Errorf("alignmentFuel = %v, want %v", alignmentFuel, tt.want)
			}
		})
	}
}
