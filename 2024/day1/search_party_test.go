package day1

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 11,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 1660292,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstParty, secondParty, err := parseSearchParties(tt.data)
			if err != nil {
				t.Fatalf("failed to parse games %v", err)
			}
			value := calcSearchPartyDistances(firstParty, secondParty)
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}

func TestSimilarity(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "AoC Example",
			data: "sample",
			want: 31,
		},
		{
			name: "AoC Task",
			data: "task",
			want: 22776016,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstParty, secondParty, err := parseSearchParties(tt.data)
			if err != nil {
				t.Fatalf("failed to parse games %v", err)
			}
			value := calcSearchPartySimilarity(firstParty, secondParty)
			if value != tt.want {
				t.Errorf("Value = %v, want %v", value, tt.want)
			}
		})
	}
}
