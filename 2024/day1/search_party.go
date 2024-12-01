package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type searchParty struct {
	locationIds []int
}

func parseSearchParties(path string) (*searchParty, *searchParty, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	var firstLocationIds []int
	var secondLocationIds []int
	raw := strings.Split(string(data), "\n")
	for _, line := range raw {
		line = standardizeSpaces(line)
		rawIds := strings.Split(line, " ")
		firstId, err := strconv.Atoi(rawIds[0])
		if err != nil {
			return nil, nil, err
		}
		firstLocationIds = append(firstLocationIds, firstId)
		secondId, err := strconv.Atoi(rawIds[1])
		if err != nil {
			return nil, nil, err
		}
		secondLocationIds = append(secondLocationIds, secondId)
	}

	return &searchParty{firstLocationIds}, &searchParty{secondLocationIds}, nil
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func calcSearchPartyDistances(firstParty, secondParty *searchParty) int {
	sum := 0
	sort.Ints(firstParty.locationIds)
	sort.Ints(secondParty.locationIds)
	for index, firstSearchLocation := range firstParty.locationIds {
		secondSearchLocation := secondParty.locationIds[index]

		sum += abs(firstSearchLocation - secondSearchLocation)
	}
	return sum
}

func calcSearchPartySimilarity(firstParty, secondParty *searchParty) int {
	sum := 0

	for _, firstSearchLocation := range firstParty.locationIds {
		occurrences := 0
		for _, secondSearchLocation := range secondParty.locationIds {
			if secondSearchLocation == firstSearchLocation {
				occurrences++
			}
		}
		sum += firstSearchLocation * occurrences
	}
	return sum
}

type searchLocation struct {
	locationId       int
	locationPosition int
}

func (l searchLocation) calcDistance(other searchLocation) int {
	return abs(l.locationPosition - other.locationPosition)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
