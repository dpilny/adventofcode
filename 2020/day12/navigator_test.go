package day12

import (
	"testing"
)

func TestNavigation(t *testing.T) {
	tests := []struct {
		name        string
		navigations string
		want        int
	}{
		{
			name:        "AoC Example part 1",
			navigations: "sample",
			want:        25,
		},
		{
			name:        "AoC Task part 1",
			navigations: "task",
			want:        2386,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			navs, err := parseNavigations(tt.navigations)
			if err != nil {
				t.Fatalf("failed to parse seatings")
			}
			navi := navigator{
				instructions: navs,
				direction:    90,
			}
			distance := navi.navigate()
			if distance != tt.want {
				t.Errorf("got = %v, want %v", distance, tt.want)
			}
		})
	}
}


func TestWaypointNavigation(t *testing.T) {
	tests := []struct {
		name        string
		navigations string
		want        int
	}{
		{
			name:        "AoC Example part 2",
			navigations: "sample",
			want:        286,
		},
		{
			name:        "AoC Task part 2",
			navigations: "task",
			want:        2386,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			navs, err := parseNavigations(tt.navigations)
			if err != nil {
				t.Fatalf("failed to parse seatings")
			}
			navi := navigator{
				instructions: navs,
				eastWaypoint: 10,
				northWaypoint: 1,
			}
			distance := navi.navigateWaypoints()
			if distance != tt.want {
				t.Errorf("got = %v, want %v", distance, tt.want)
			}
		})
	}
}
