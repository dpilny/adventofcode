package day4

import "testing"

func TestBingoFirstWin(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 4512,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 12796,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bingoGame, err := GetGame(tt.data)
			if err != nil {
				t.Fatalf("failed to parse bingoGame %v", err)
			}
			resultValue := bingoGame.firstWin()
			if err != nil {
				t.Fatalf("failed get resultValue %v", err)
			}
			if resultValue != tt.want {
				t.Errorf("resultValue = %v, want %v", resultValue, tt.want)
			}
		})
	}
}

func TestBingoLastWin(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 1924,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 18063,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bingoGame, err := GetGame(tt.data)
			if err != nil {
				t.Fatalf("failed to parse bingoGame %v", err)
			}
			resultValue := bingoGame.lastWin()
			if err != nil {
				t.Fatalf("failed get resultValue %v", err)
			}
			if resultValue != tt.want {
				t.Errorf("resultValue = %v, want %v", resultValue, tt.want)
			}
		})
	}
}
