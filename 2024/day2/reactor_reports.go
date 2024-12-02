package day2

import (
	"os"
	"strconv"
	"strings"
)

type report struct {
	levels []int
}

type reports []report

func parseReports(path string) (*reports, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var r reports
	raw := strings.Split(string(data), "\n")
	for _, line := range raw {
		var levels []int
		rawIds := strings.Split(line, " ")
		for _, lvl := range rawIds {
			val, err := strconv.Atoi(lvl)
			if err != nil {
				return nil, err
			}
			levels = append(levels, val)
		}
		r = append(r, report{levels: levels})
	}
	return &r, nil
}

func isReportSafe(levels []int) bool {
	prev := -1
	increasing := false
	for idx, level := range levels {
		if idx == 1 {
			increasing = level > prev
		}
		if idx > 0 {
			if increasing && level < prev {
				return false
			}
			if !increasing && level > prev {
				return false
			}
			diff := abs(prev - level)
			if diff > 3 || diff == 0 {
				return false
			}
		}
		prev = level
	}
	return true
}

func (r *reports) calculateSafeReports() int {
	safeReports := 0
	for _, report := range *r {
		if isReportSafe(report.levels) {
			safeReports++
		}
	}
	return safeReports
}

func remove(slice []int, s int) []int {
	cpy := append([]int(nil), slice...)
	return append(cpy[:s], cpy[s+1:]...)
}

func (r *reports) calculateSafeDampenedReports() int {
	safeReports := 0
exit:
	for _, report := range *r {
		if isReportSafe(report.levels) {
			safeReports++
			continue
		} else {
			for idx, _ := range report.levels {
				dampenedLevels := remove(report.levels, idx)
				if isReportSafe(dampenedLevels) {
					safeReports++
					continue exit
				}
			}
		}
	}
	return safeReports
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
