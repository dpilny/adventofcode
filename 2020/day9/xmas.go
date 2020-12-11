package day9

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Xmas struct {
	PreambleLength int
	Data           []int64
	// contains possible sums from current preamble window as keys in map
	// values contains a slice of indices which contribute to the sum
	sums map[int64][]sumOf

	failedNumber int64
}

type sumOf struct {
	a int64
	b int64
}

func Get(preambleLength int, data []int64) Xmas {
	return Xmas{
		PreambleLength: preambleLength,
		Data:           data,
		sums:           map[int64][]sumOf{},
	}
}

func parseData(dataPath string) ([]int64, error) {
	content, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(content), "\n")
	var data []int64
	for _, rawLine := range raw {
		value, err := strconv.ParseInt(rawLine, 10, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}
	return data, nil
}

func (x Xmas) findContiguousSum() int64 {
	invalidVal, index := x.findInvalid()

	for i := 0; i < index; i++ {
		var sum int64 = 0
		var summands []int64
		for j := i; j < index; j++ {
			sum += x.Data[j]
			summands = append(summands, x.Data[j])
			if sum == invalidVal {
				min := findMin(summands)
				max := findMax(summands)
				log.Printf("found contiguous sum for val: %v min: %v max: %v summands: %v+", invalidVal, min, max, summands)
				return min + max
			} else if sum > invalidVal {
				break
			}
		}
	}
	return -1
}

func findMin(data []int64) int64 {
	var min int64
	for i, v := range data {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

func findMax(data []int64) int64 {
	var max int64
	for i, v := range data {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

// returns invalid value from the sequence with its index, returns -1 if non was found
func (x Xmas) findInvalid() (int64, int) {
	for i, val := range x.Data {
		//x.updateSums(i)
		if i < x.PreambleLength {
			continue
		}
		possibleSums := map[int64]sumOf{}

		for k := i - x.PreambleLength; k < i; k++ {
			for j := k + 1; j < i; j++ {
				sum := x.Data[k] + x.Data[j]
				possibleSums[sum] = sumOf{
					a: x.Data[k],
					b: x.Data[j],
				}
			}
		}

		log.Printf("possible sums for index and value: %v, %v, %+v", i, val, possibleSums)

		_, ok := possibleSums[val]
		if !ok {
			return val, i
		}

	}
	return -1, -1
}

func (x Xmas) updateSums(index int) {
	startIndex := index - x.PreambleLength
	if startIndex < 0 {
		return
	}
	//for k := startIndex; k < index-1; k++ {
	//	for j := k + 1; j < index; j++ {
	//		sum := x.Data[k] + x.Data[j]
	//		val, ok := x.sums[sum]
	//		if ok {
	//			x.sums[sum] = append(val, sumOf{
	//				a: k,
	//				b: j,
	//			})
	//		} else {
	//			x.sums[sum] = []sumOf{{
	//				a: k,
	//				b: j,
	//			}}
	//		}
	//	}
	//}
	removeIndex := startIndex - 1
	if removeIndex >= 0 {
		// remove all sums which were calculated by the value from the old index
		for sum, sumOfs := range x.sums {
			i := 0
			//for _, sumOf := range sumOfs {
			//if sumOf.a == removeIndex || sumOf.b == removeIndex {
			//	copy and increment index
			//sumOfs[i] = sumOf
			//i++
			//}
			//}
			// Prevent memory leak by erasing truncated values
			sumOfs = sumOfs[:i]
			if len(sumOfs) == 0 {
				delete(x.sums, sum)
			} else {
				x.sums[sum] = sumOfs
			}
		}
	}
}
