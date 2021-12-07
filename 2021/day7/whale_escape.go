package day7

import (
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseCrabAlignment(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	rawPopulation := strings.Split(string(data), ",")
	var positions []int
	for _, rawInput := range rawPopulation {
		position, err := strconv.Atoi(rawInput)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func calculateAlignLeastFuel(positions []int) int {
	mean := calcMean(positions)
	neededFuel := 0
	for _, v := range positions {
		neededFuel += int(math.Abs(float64(v - mean)))
	}
	return neededFuel
}

func calcMean(values []int) int {
	sort.Ints(values)
	var mean int
	posCount := len(values)
	if posCount%2 == 0 {
		mean = (values[posCount/2] + values[posCount/2-1]) / 2
	} else {
		mean = values[posCount/2-1]
	}
	return mean
}

func calcAvg(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	avg := int(math.Round(float64(sum) / float64(len(values))))
	return avg
}

func calculateAlignBruteForce(positions []int) int {
	sort.Ints(positions)
	lowerBound := positions[0]
	upperBound := positions[len(positions)-1]
	lowestFuel := math.MaxInt64
	for i := lowerBound; i <= upperBound; i++ {
		neededFuel := 0
		for _, v := range positions {
			neededFuel += sumTill(int(math.Abs(float64(v - i))))
		}
		if neededFuel < lowestFuel {
			lowestFuel = neededFuel
		}
	}
	return lowestFuel
}

func sumTill(val int) int {
	sum := 0
	for i := 1; i <= val; i++ {
		sum += i
	}
	return sum
}
