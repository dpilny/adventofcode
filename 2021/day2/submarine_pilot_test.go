package day2

import "testing"

func TestMultipliedPosition(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 150,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 1962940,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commands, err := parseCommands(tt.data)
			if err != nil {
				t.Fatalf("failed to parse commands")
			}
			product := calcMultipliedPosition(commands)
			if product != tt.want {
				t.Errorf("Product = %v, want %v", product, tt.want)
			}
		})
	}
}

func TestMultipliedPositionAimed(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 900,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 1813664422,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commands, err := parseCommands(tt.data)
			if err != nil {
				t.Fatalf("failed to parse commands")
			}
			product := calcMultipliedPositionAimed(commands)
			if product != tt.want {
				t.Errorf("Product = %v, want %v", product, tt.want)
			}
		})
	}
}


