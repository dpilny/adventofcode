package day3

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type DiagnosticData struct {
	length     int
	width      int
	binaryData [][]bool
}

func parseDiagnosticInput(path string) (*DiagnosticData, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(data), "\n")

	var binaryData [][]bool
	width := 0
	for _, rawLine := range raw {
		var rowData []bool
		for _, char := range rawLine {
			val := false
			if char == '1' {
				val = true
			}
			rowData = append(rowData, val)
		}
		if width == 0 {
			width = len(rowData)
		}
		binaryData = append(binaryData, rowData)
	}

	return &DiagnosticData{
		length:     len(raw),
		width:      width,
		binaryData: binaryData,
	}, nil
}

func calculatePowerConsumption(data *DiagnosticData) (int, error) {
	var gammaBuilder strings.Builder
	var epsilonBuilder strings.Builder
	mostCommonThreshold := data.length / 2
	for x := 0; x < data.width; x++ {
		trueCount := 0
		for y := 0; y < data.length; y++ {
			if data.binaryData[y][x] {
				trueCount++
			}
		}
		if trueCount > mostCommonThreshold {
			gammaBuilder.WriteRune('1')
			epsilonBuilder.WriteRune('0')
		} else {
			gammaBuilder.WriteRune('0')
			epsilonBuilder.WriteRune('1')
		}
	}
	gamma, err := strconv.ParseInt(gammaBuilder.String(), 2, 64)
	if err != nil {
		return -1, err
	}
	epsilon, err := strconv.ParseInt(epsilonBuilder.String(), 2, 64)
	if err != nil {
		return -1, err
	}
	return int(gamma * epsilon), nil
}

func removeIndexFromSlice(s [][]bool, index int) [][]bool {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeIndicesFromSlice(indices []int, data [][]bool) [][]bool {
	for i := len(indices) - 1; i >= 0; i-- {
		indexToRemove := indices[i]
		data = removeIndexFromSlice(data, indexToRemove)
	}
	return data
}

func boolArrayToInt(data []bool) int {
	val := 0
	for i, datum := range data {
		exp := len(data) - (i + 1)
		if datum {
			val += int(math.Pow(2, float64(exp)))
		}
	}
	return val
}

func copyData(src [][]bool, dst [][]bool) {
	for i := range src {
		dst[i] = make([]bool, len(src[i]))
		copy(dst[i], src[i])
	}
}

func countTrueEntries(x int, data [][]bool) ([]int, []int) {
	var zeroRows []int
	var oneRows []int
	for y := 0; y < len(data); y++ {
		if data[y][x] {
			oneRows = append(oneRows, y)
		} else {
			zeroRows = append(zeroRows, y)
		}
	}
	return zeroRows, oneRows
}

func calculateRating(data [][]bool, prioritizeOne bool) int {
	copiedData := make([][]bool, len(data))
	copyData(data, copiedData)
	for hasEntries := true; hasEntries; hasEntries = len(copiedData) > 1 {
		for x := 0; x < len(data); x++ {
			zeroRows, oneRows := countTrueEntries(x, copiedData)
			falseCount := len(zeroRows)
			trueCount := len(oneRows)

			if trueCount >= falseCount {
				copiedData = removeIndicesFromSlice(zeroRows, copiedData)
			} else {
				copiedData = removeIndicesFromSlice(oneRows, copiedData)
			}
			if len(copiedData) == 1 {
				break
			}
		}
	}
	return boolArrayToInt(copiedData[0])
}

func calculateLifeSupportRating(data *DiagnosticData) (int, error) {
	// oxygen rating
	oxygenEntries := make([][]bool, len(data.binaryData))
	copyData(data.binaryData, oxygenEntries)
	for hasEntries := true; hasEntries; hasEntries = len(oxygenEntries) > 1 {
		for x := 0; x < data.width; x++ {
			zeroRows, oneRows := countTrueEntries(x, oxygenEntries)
			falseCount := len(zeroRows)
			trueCount := len(oneRows)

			if trueCount >= falseCount {
				oxygenEntries = removeIndicesFromSlice(zeroRows, oxygenEntries)
			} else {
				oxygenEntries = removeIndicesFromSlice(oneRows, oxygenEntries)
			}
			if len(oxygenEntries) == 1 {
				break
			}
		}
	}
	oxygenRating := boolArrayToInt(oxygenEntries[0])

	// co2 scrubber
	co2ScrubberEntries := make([][]bool, len(data.binaryData))
	copyData(data.binaryData, co2ScrubberEntries)
	for hasEntries := true; hasEntries; hasEntries = len(co2ScrubberEntries) > 1 {
		for x := 0; x < data.width; x++ {
			zeroRows, oneRows := countTrueEntries(x, co2ScrubberEntries)
			trueCount := len(oneRows)
			falseCount := len(zeroRows)

			if trueCount >= falseCount {
				co2ScrubberEntries = removeIndicesFromSlice(oneRows, co2ScrubberEntries)
			} else {
				co2ScrubberEntries = removeIndicesFromSlice(zeroRows, co2ScrubberEntries)
			}
			if len(co2ScrubberEntries) == 1 {
				break
			}
		}
	}
	co2ScrubberRating := boolArrayToInt(co2ScrubberEntries[0])

	return oxygenRating * co2ScrubberRating, nil
}
