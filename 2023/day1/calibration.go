package day1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type numberRepresentation struct {
	runes       []rune
	replaceWith rune
	runIndex    int
}

func replaceWrittenNumbersInLine(line string) string {
	fmt.Printf("in: %v\n", line)
	reps := []numberRepresentation{
		{
			runes:       []rune("one"),
			replaceWith: '1',
			runIndex:    0,
		},
		{
			runes:       []rune("two"),
			replaceWith: '2',
			runIndex:    0,
		},
		{
			runes:       []rune("three"),
			replaceWith: '3',
			runIndex:    0,
		},
		{
			runes:       []rune("four"),
			replaceWith: '4',
			runIndex:    0,
		},
		{
			runes:       []rune("five"),
			replaceWith: '5',
			runIndex:    0,
		},
		{
			runes:       []rune("six"),
			replaceWith: '6',
			runIndex:    0,
		},
		{
			runes:       []rune("seven"),
			replaceWith: '7',
			runIndex:    0,
		},
		{
			runes:       []rune("eight"),
			replaceWith: '8',
			runIndex:    0,
		},
		{
			runes:       []rune("nine"),
			replaceWith: '9',
			runIndex:    0,
		},
	}
	reachedEnd := false
	breakLoop := false
	skipUntil := 0
	for !reachedEnd {
		for pos, lineRune := range line {
			if skipUntil >= len(line) {
				reachedEnd = true
			}
			if pos < skipUntil {
				continue
			}
			for posRep, rep := range reps {
				if lineRune == rep.runes[rep.runIndex] {
					reps[posRep].runIndex = rep.runIndex + 1
					if reps[posRep].runIndex == len(rep.runes) {
						// match - replace string
						line = strings.Replace(line, string(rep.runes), string(rep.replaceWith), 1)
						line = strings.Replace(line, string(rep.runes[1:]), string(rep.replaceWith), 1)
						reps[posRep].runIndex = 0
						skipUntil = pos - len(rep.runes) + 2
						breakLoop = true
					}
				} else {
					reps[posRep].runIndex = 0
				}
			}
			if breakLoop {
				breakLoop = false
				break
			}
			if pos == len(line)-1 {
				reachedEnd = true
			}
		}
	}
	fmt.Printf("out: %v\n", line)
	return line
}

func getCalibrationValue(path string, complex bool) (int, error) {
	sum := 0
	values, err := parseCalibrationFile(path, complex)
	if err != nil {
		return -1, err
	}
	for _, value := range values {
		sum += value
	}
	return sum, err
}

func parseCalibrationFile(path string, complex bool) ([]int, error) {
	r, err := regexp.Compile("\\D+")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	stringValue := strings.ToLower(string(data))
	raw := strings.Split(stringValue, "\n")
	var values []int
	for _, rawLine := range raw {
		if complex {
			rawLine = replaceWrittenNumbersInLine(rawLine)
		}
		digits := r.ReplaceAllString(rawLine, "")
		rawValue := string(digits[0]) + string(digits[len(digits)-1])
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}
