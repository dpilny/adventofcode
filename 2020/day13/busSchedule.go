package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func process2nNote(note string) int {
	parts := strings.Split(note, "\n")
	var busIds []int
	for _, idValue := range strings.Split(parts[1], ",") {
		if idValue == "x" {
			busIds = append(busIds, 0)
			continue
		}
		busId, err := strconv.Atoi(idValue)
		if err != nil {
			panic(fmt.Errorf("invalid bus id %v\n", busId))
		}
		busIds = append(busIds, busId)
	}
	return calcSubsequentDepartures(busIds)
}

func calcSubsequentDepartures(busIds []int) int {
	index := 0
	timestamp := 0

	busMultiplier := 1
	nextBusMultiplier := 1
	for {

		bus := busIds[index]
		nextBus := busIds[index+1]

		// busId represents wildcard busId
		if bus == 0 || nextBus == 0 {
			index++
		}

		// find subsequent value for bus & nextBus
		for {
			busValue := bus * busMultiplier
			nextBusValue := nextBus * nextBusMultiplier
			if busValue+1 == nextBusValue {
				if index == 0 {
					timestamp = busValue
				}
				index++
				break
			} else if busValue > nextBusValue {
				nextBusMultiplier++
			} else if nextBusValue > busValue && index == 0 {
				busMultiplier++
			} else {
				index = 0
			}
		}
		index++
		if index == len(busIds)-1 {
			break
		}
	}

	return timestamp
}

func processNote(note string) int {
	parts := strings.Split(note, "\n")
	departureAt, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Errorf("failed to parse departure at value %v\n", parts[0]))
	}
	var busIds []int
	for _, idValue := range strings.Split(parts[1], ",") {
		if idValue == "x" {
			continue
		}
		busId, err := strconv.Atoi(idValue)
		if err != nil {
			panic(fmt.Errorf("invalid bus id %v\n", busId))
		}
		busIds = append(busIds, busId)
	}
	return calcDepartureTime(departureAt, busIds)
}

func calcDepartureTime(departureAt int, busIds []int) int {
	minWaitTime := math.MaxInt64
	minWaitTimeBusId := 0
	for _, busId := range busIds {
		x := departureAt / busId
		y := x * busId
		if y < departureAt {
			y += busId
		}
		waitTime := y - departureAt
		if waitTime < minWaitTime {
			minWaitTime = waitTime
			minWaitTimeBusId = busId
		}
	}
	return minWaitTimeBusId * minWaitTime
}
