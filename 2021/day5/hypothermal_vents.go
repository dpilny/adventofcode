package day5

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type HypoMap struct {
	lines []Line
}

type Line struct {
	origin Coordinate
	dest   Coordinate
}

type Coordinate struct {
	x, y int
}

func parseHypoLines(path string) (*HypoMap, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(data), "\n")

	var lines []Line
	for _, rawLine := range raw {
		fields := strings.Fields(rawLine)
		x1, y1, err := parseCoordinates(fields[0])
		if err != nil {
			return nil, err
		}
		x2, y2, err := parseCoordinates(fields[2])
		if err != nil {
			return nil, err
		}
		lines = append(lines, Line{
			origin: Coordinate{
				x: x1,
				y: y1,
			},
			dest: Coordinate{
				x: x2,
				y: y2,
			},
		})
	}

	return &HypoMap{
		lines: lines,
	}, nil
}

func parseCoordinates(data string) (int, int, error) {
	parts := strings.Split(data, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1, -1, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, -1, err
	}
	return x, y, nil
}

func (l *Line) isDiagonal() bool {
	return !l.isHorizontal() && !l.isVertical()
}

func (l *Line) isHorizontal() bool {
	return l.origin.y == l.dest.y
}

func (l *Line) isVertical() bool {
	return l.origin.x == l.dest.x
}

func (l Line) getNonDiagonalBounds() (int, int, int, int) {
	var lowerX, upperX, lowerY, upperY int
	if l.origin.x < l.dest.x {
		lowerX = l.origin.x
		upperX = l.dest.x
	} else {
		lowerX = l.dest.x
		upperX = l.origin.x
	}
	if l.origin.y < l.dest.y {
		lowerY = l.origin.y
		upperY = l.dest.y
	} else {
		lowerY = l.dest.y
		upperY = l.origin.y
	}
	return lowerX, upperX, lowerY, upperY
}
func (l Line) getDiagonalBounds() (Coordinate, Coordinate) {
	if l.origin.x < l.dest.x {
		return l.origin, l.dest
	} else {
		return l.dest, l.origin
	}
}

func (l *Line) getPathCoordinates() []Coordinate {
	var coordinates []Coordinate

	if l.isDiagonal() {
		lowerXCord, upperXCord := l.getDiagonalBounds()
		y := lowerXCord.y
		for x := lowerXCord.x; x <= upperXCord.x; x++ {
			increment := lowerXCord.y < upperXCord.y
			coordinates = append(coordinates, Coordinate{
				x: x,
				y: y,
			})
			if increment {
				y++
			} else {
				y--
			}
		}
	} else {
		lowerX, upperX, lowerY, upperY := l.getNonDiagonalBounds()
		if l.isHorizontal() {
			for x := lowerX; x <= upperX; x++ {
				coordinates = append(coordinates, Coordinate{
					x: x,
					y: l.origin.y,
				})
			}
		} else if l.isVertical() {
			for y := lowerY; y <= upperY; y++ {
				coordinates = append(coordinates, Coordinate{
					x: l.origin.x,
					y: y,
				})
			}
		}
	}

	return coordinates
}

func (m HypoMap) countOverlaps(ignoreDiagonal bool) int {
	occurrenceMap := map[Coordinate][]Line{}
	for _, line := range m.lines {
		if ignoreDiagonal && line.isDiagonal() {
			continue
		}
		cords := line.getPathCoordinates()
		for _, cord := range cords {
			if val, ok := occurrenceMap[cord]; ok {
				occurrenceMap[cord] = append(val, line)
			} else {
				occurrenceMap[cord] = []Line{line}
			}
		}
	}

	overlapCount := 0
	for _, lines := range occurrenceMap {
		if len(lines) > 1 {
			overlapCount++
		}
	}

	return overlapCount
}
