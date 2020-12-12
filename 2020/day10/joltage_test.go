package day10

import (
	"log"
	"testing"
)

func TestJoltageChain(t *testing.T) {
	tests := []struct {
		name     string
		adapters string
		want     int
	}{
		{
			name:     "AoC Example",
			adapters: "sample",
			want:     220,
		},
		{
			name:     "AoC Task",
			adapters: "task",
			want:     2812,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := parseAdapters(tt.adapters)
			if err != nil {
				t.Fatalf("failed to parse instructions")
			}
			oneDiffs, twoDiffs, threeDiffs := getJoltageChainDiffs(data)
			log.Printf("got %v oneDiffs, %v twoDiffs, and %v threeDiffs", oneDiffs, twoDiffs, threeDiffs)
			if oneDiffs*threeDiffs != tt.want {
				t.Errorf("got = %v, want %v", oneDiffs*threeDiffs, tt.want)
			}
		})
	}
}

func TestDistincJoltageChainCount(t *testing.T) {
	tests := []struct {
		name     string
		adapters string
		want     int
	}{
		{
			name:     "AoC Example",
			adapters: "sample",
			want:     19208,
		},
		//{
		//	name:     "AoC Task",
		//	adapters: "task",
		//	want:     2812,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := parseAdapters(tt.adapters)
			if err != nil {
				t.Fatalf("failed to parse instructions")
			}
			chainCount := getDistinctJoltageChainCount(data)
			log.Printf("got %v distincs chains", chainCount)
			if chainCount != tt.want {
				t.Errorf("got = %v, want %v", chainCount, tt.want)
			}
		})
	}
}
