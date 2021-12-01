package day13

import "testing"

func TestBusSchedule(t *testing.T) {
	tests := []struct {
		name        string
		busSchedule string
		want        int
	}{
		{
			name:        "AoC Example part 1",
			busSchedule: "939\n7,13,x,x,59,x,31,19",
			want:        295,
		},
		{
			name:        "AoC Task part 1",
			busSchedule: "1008169\n29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,653,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,x,x,823,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19",
			want:        4938,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := processNote(tt.busSchedule)
			if value != tt.want {
				t.Errorf("got = %v, want %v", value, tt.want)
			}
		})
	}
}
