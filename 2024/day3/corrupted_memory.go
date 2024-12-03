package day3

import (
	"os"
	"regexp"
	"strconv"
)

type memory string

func parseMemory(path string) (*memory, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	mem := memory(data)
	return &mem, nil
}

func (m memory) calculateMultiplicationSum() int {
	r := regexp.MustCompile("mul\\((?P<factor1>\\d{1,3}),(?P<factor2>\\d{1,3})\\)")
	matches := r.FindAllStringSubmatch(string(m), -1)
	sum := 0
	for _, match := range matches {
		factor1, _ := strconv.Atoi(match[1])
		factor2, _ := strconv.Atoi(match[2])
		sum += factor1 * factor2
	}
	return sum
}

func (m memory) calculateSanitizedMultiplicationSum() int {
	r := regexp.MustCompile("(don't\\(\\))|(do\\(\\))|(mul\\((?P<factor1>\\d{1,3}),(?P<factor2>\\d{1,3})\\))")
	matches := r.FindAllStringSubmatch(string(m), -1)
	sum := 0
	active := true
	for _, match := range matches {
		if match[0] == "do()" {
			active = true
		} else if match[0] == "don't()" {
			active = false
		} else if active {
			factor1, _ := strconv.Atoi(match[4])
			factor2, _ := strconv.Atoi(match[5])
			sum += factor1 * factor2
		}
	}
	return sum
}
