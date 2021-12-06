package day6

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type LanternFishPopulation struct {
	fish []int
}

type LanternFish struct {
	daysTillReproduction int
}

func parseInitialPopulation(path string) (*LanternFishPopulation, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	rawPopulation := strings.Split(string(data), ",")
	var population []int
	for _, rawInput := range rawPopulation {
		daysTillReproduction, err := strconv.Atoi(rawInput)
		if err != nil {
			return nil, err
		}
		population = append(population, daysTillReproduction)
	}
	return &LanternFishPopulation{fish: population}, nil
}

func (p LanternFishPopulation) countPopulationAfterDays(days int) int {
	return p.countByMap(days)
}

func (p LanternFishPopulation) countByMap(days int) int {
	popMap := map[int]int{}
	for _, age := range p.fish {
		if val, ok := popMap[age]; ok {
			popMap[age] = val + 1
		} else {
			popMap[age] = 1
		}
	}

	for day := 0; day < days; day++ {
		nextPop := map[int]int{}
		for age, count := range popMap {
			newAge := age - 1
			if newAge == -1 {
				newAge = 6
				nextPop[8] = count
			}
			if val, ok := nextPop[newAge]; ok {
				nextPop[newAge] = val + count
			} else {
				nextPop[newAge] = count
			}
		}
		popMap = nextPop
	}
	population := 0
	for _, v := range popMap {
		population += v
	}

	return population
}

func (p LanternFishPopulation) countByArray(days int) int {
	for day := 0; day < days; day++ {
		popSize := len(p.fish)
		for i := 0; i < popSize; i++ {
			newAge := p.fish[i] - 1
			if newAge == -1 {
				newAge = 6
				p.fish = append(p.fish, 8)
			}
			p.fish[i] = newAge
		}
		fmt.Printf("at day %v\n", day+1)
	}
	return len(p.fish)
}
