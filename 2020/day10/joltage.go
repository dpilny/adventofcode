package day10

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseAdapters(bagPath string) ([]int, error) {
	content, err := ioutil.ReadFile(bagPath)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(content), "\n")
	var data []int
	for _, rawLine := range raw {
		value, err := strconv.Atoi(rawLine)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}
	return data, nil
}

func getJoltageChainDiffs(adapters []int) (int, int, int) {
	joltDiffs := map[int]int{}

	sort.Ints(adapters)

	curJoltage := 0
	for _, val := range adapters {
		diff := val - curJoltage
		if diff > 3 {
			panic(fmt.Errorf("invalid joltage diff %v", diff))
		}
		joltDiffs[diff] = joltDiffs[diff] + 1
		curJoltage = val
	}
	// add one final three jolt diff for the final output
	joltDiffs[3] = joltDiffs[3] + 1
	return joltDiffs[1], joltDiffs[2], joltDiffs[3]
}

func getDistinctJoltageChainCount(adapters []int) int {
	sort.Ints(adapters)

	joltageChains := [][]int{
		{0},
	}
	//1259712
	for _, adapter := range adapters {
		for i, chain := range joltageChains {
			for k, value := range chain {
				if adapter-value <= 3 {
					joltageChains[i] = append(chain, adapter)
					if k < len(chain) {
						// copy current chain until k
						cp := make([]int, len(joltageChains[i][:k]))
						copy(cp, joltageChains[i][:k])
						joltageChains = append(joltageChains, cp)
					}
				} else {

				}
			}
		}
	}
	log.Printf("got joltage chains: %v+", joltageChains)
	return len(joltageChains)
}
