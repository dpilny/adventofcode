package day1

import "testing"

func TestSweepIncrease(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 7,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 1167,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			measurements, err := parseMeasurements(tt.data)
			if err != nil {
				t.Fatalf("failed to parse measurements")
			}
			incs := countIncreases(measurements)
			if incs != tt.want {
				t.Errorf("Incs = %v, want %v", incs, tt.want)
			}
		})
	}
}

func TestSweepWindowIncrease(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 5,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 1130,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			measurements, err := parseMeasurements(tt.data)
			if err != nil {
				t.Fatalf("failed to parse measurements")
			}

			windows := getSlidingWindows(measurements)

			incs := countIncreases(windows)
			if incs != tt.want {
				t.Errorf("Incs = %v, want %v", incs, tt.want)
			}
		})
	}
}
