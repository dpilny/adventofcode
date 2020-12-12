package day9

import "testing"

func TestFindInvalid(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		preambleLength int
		want           int64
	}{
		{
			name:           "AoC Example",
			code:           "sample",
			want:           127,
			preambleLength: 5,
		},
		{
			name:           "AoC Task",
			code:           "task",
			want:           85848519,
			preambleLength: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := parseData(tt.code)
			if err != nil {
				t.Fatalf("failed to parse instructions")
			}
			xmas := Get(tt.preambleLength, data)
			failed, _ := xmas.findInvalid()
			if failed != tt.want {
				t.Errorf("failed at = %v, want %v", failed, tt.want)
			}
		})
	}
}


func TestFindContiguousSet(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		preambleLength int
		want           int64
	}{
		{
			name:           "AoC Example",
			code:           "sample",
			want:           62,
			preambleLength: 5,
		},
		{
			name:           "AoC Task",
			code:           "task",
			want:           13414198,
			preambleLength: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := parseData(tt.code)
			if err != nil {
				t.Fatalf("failed to parse instructions")
			}
			xmas := Get(tt.preambleLength, data)
			failed := xmas.findContiguousSum()
			if failed != tt.want {
				t.Errorf("failed at = %v, want %v", failed, tt.want)
			}
		})
	}
}