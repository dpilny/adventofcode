package day1

import "testing"

func TestLineReplace(t *testing.T) {
	line := "eightwothree"
	replaced := replaceWrittenNumbersInLine(line)
	if replaced != "823" {
		t.Errorf("Value = %v, want %v", replaced, "8wo3")
	}

}

func TestSimpleParseCalibration(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 142,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 54630,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := getCalibrationValue(tt.data, false)
			if err != nil {
				t.Fatalf("failed to parse measurements %v", err)
			}
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}

func TestComplexParseCalibration(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample_2",
			want: 281,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 54782,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := getCalibrationValue(tt.data, true)
			if err != nil {
				t.Fatalf("failed to parse measurements %v", err)
			}
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}
