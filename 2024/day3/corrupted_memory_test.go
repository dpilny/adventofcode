package day3

import "testing"

func TestMemoryCorruption(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample_1",
			want: 161,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 166357705,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memory, err := parseMemory(tt.data)
			if err != nil {
				t.Fatalf("failed to parse games %v", err)
			}
			value := memory.calculateMultiplicationSum()
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}

func TestSanitziedMemoryCorruption(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample_2",
			want: 48,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 88811886,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memory, err := parseMemory(tt.data)
			if err != nil {
				t.Fatalf("failed to parse games %v", err)
			}
			value := memory.calculateSanitizedMultiplicationSum()
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}
