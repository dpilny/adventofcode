package day11

import (
	"log"
	"testing"
)

func TestSeatSettling(t *testing.T) {
	tests := []struct {
		name            string
		seating         string
		directNeighbors bool
		want            int
	}{
		{
			name:            "AoC Example part 1",
			seating:         "sample",
			directNeighbors: true,
			want:            37,
		},
		{
			name:            "AoC Task part 1",
			seating:         "task",
			directNeighbors: true,
			want:            2386,
		},
		{
			name:            "AoC Example part 2",
			seating:         "sample",
			directNeighbors: false,
			want:            26,
		},
		{
			name:            "AoC Example part 2",
			seating:         "task",
			directNeighbors: false,
			want:            2091,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layout, err := parse(tt.seating, tt.directNeighbors)
			if err != nil {
				t.Fatalf("failed to parse seatings")
			}
			roundCount, occupiedSeats := layout.letPeopleSit()
			log.Printf("seating settled after %v rounds with %v seats occupied", roundCount, occupiedSeats)
			if occupiedSeats != tt.want {
				t.Errorf("got = %v, want %v", occupiedSeats, tt.want)
			}
		})
	}
}
