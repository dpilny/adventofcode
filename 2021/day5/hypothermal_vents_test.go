package day5

import "testing"

func TestLineOverlapNonDiagonal(t *testing.T) {
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
			want: 5698,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hypoLines, err := parseHypoLines(tt.data)
			if err != nil {
				t.Fatalf("failed to parse hypoLines %v", err)
			}
			resultValue := hypoLines.countOverlaps(true)
			if err != nil {
				t.Fatalf("failed get overlaps %v", err)
			}
			if resultValue != tt.want {
				t.Errorf("overlaps = %v, want %v", resultValue, tt.want)
			}
		})
	}
}

func TestLineOverlap(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 12,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 15463,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hypoLines, err := parseHypoLines(tt.data)
			if err != nil {
				t.Fatalf("failed to parse hypoLines %v", err)
			}
			resultValue := hypoLines.countOverlaps(false)
			if err != nil {
				t.Fatalf("failed get overlaps %v", err)
			}
			if resultValue != tt.want {
				t.Errorf("overlaps = %v, want %v", resultValue, tt.want)
			}
		})
	}
}
