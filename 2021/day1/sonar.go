package day1

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func parseMeasurements(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(data), "\n")
	var lines []int
	for _, rawLine := range raw {
		measurement, err := strconv.Atoi(rawLine)
		if err != nil {
			return nil, err
		}
		lines = append(lines, measurement)
	}
	return lines, nil
}

func countIncreases(measurements []int) int {
	count := 0

	prevValue := math.MaxInt64

	for _, measurement := range measurements {
		if measurement > prevValue {
			count++
		}
		prevValue = measurement
	}

	return count
}

func getSlidingWindows(measurements []int) []int {
	var windowSums []int
	for i := 3; i <= len(measurements); i++ {
		possibleWindow := measurements[i-3 : i]
		sum := 0
		for _, measurement := range possibleWindow {
			sum += measurement
		}
		windowSums = append(windowSums, sum)
	}
	return windowSums
}
