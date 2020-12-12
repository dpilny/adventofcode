package day12

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	north = iota + 1
	south
	east
	west
	left
	right
	forward
)

type navIns struct {
	action byte
	value  int
}

func parseNavigations(navPath string) ([]navIns, error) {
	content, err := ioutil.ReadFile(navPath)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(content), "\n")
	var instructions []navIns
	for _, rawLine := range raw {
		var action byte
		switch string(rawLine[0]) {
		case "N":
			action = north
		case "S":
			action = south
		case "W":
			action = west
		case "E":
			action = east
		case "L":
			action = left
		case "R":
			action = right
		case "F":
			action = forward
		default:
			panic(fmt.Errorf("unknown navigation command typs %v", string(rawLine[0])))
		}
		value, err := strconv.Atoi(rawLine[1:])
		if err != nil {
			panic(fmt.Errorf("invalid value %v\n", value))
		}
		instructions = append(instructions, navIns{
			action: action,
			value:  value,
		})
	}
	return instructions, nil
}

type navigator struct {
	instructions []navIns

	eastFerry    int
	northFerry   int
	direction    int

	northWaypoint int
	eastWaypoint  int
}

func (n *navigator) navigateWaypoints() int {
	for _, ins := range n.instructions {
		switch ins.action {
		case north:
			n.northWaypoint += ins.value
		case south:
			n.northWaypoint -= ins.value
		case east:
			n.eastWaypoint += ins.value
		case west:
			n.eastWaypoint -= ins.value
		case left:
			switch ins.value {
			case 90:
				// north will be west (east + -1)
				// east will be north
				tmp := n.eastWaypoint
				n.eastWaypoint = n.northWaypoint * -1
				n.northWaypoint = tmp
			case 180:
				// north will be south (north * -1)
				// east will be west (east * -1)
				n.northWaypoint *= -1
				n.eastWaypoint *= -1
			case 270:
				// north will be east
				// east will be south (north * -1)
				tmp := n.eastWaypoint * - 1
				n.eastWaypoint = n.northWaypoint
				n.northWaypoint = tmp
			}
		case right:
			switch ins.value {
			case 90:
				// east will be north
				// north will be east * -1
				tmp := n.eastWaypoint * -1
				n.eastWaypoint = n.northWaypoint
				n.northWaypoint = tmp
			case 180:
				// north will be south (north * -1)
				// east will be west (east * -1)
				n.northWaypoint *= -1
				n.eastWaypoint *= -1
			case 270:
				// north will be west
				// east will north
				tmp := n.eastWaypoint
				n.eastWaypoint = n.northWaypoint * -1
				n.northWaypoint = tmp
			}
		case forward:
			n.northFerry += ins.value * n.northWaypoint
			n.eastFerry += ins.value * n.eastWaypoint
		}
	}
	if n.eastFerry < 0 {
		n.eastFerry *= -1
	}
	if n.northFerry < 0 {
		n.northFerry *= -1
	}
	return n.northFerry + n.eastFerry

}

// navigates the ship by the given instructions
// returns the moved manhattan distance
func (n *navigator) navigate() int {
	for _, ins := range n.instructions {
		switch ins.action {
		case north:
			n.northFerry += ins.value
		case south:
			n.northFerry -= ins.value
		case east:
			n.eastFerry += ins.value
		case west:
			n.eastFerry -= ins.value
		case left:
			n.direction = mod(n.direction-ins.value, 360)
		case right:
			n.direction = mod(n.direction+ins.value, 360)
		case forward:
			switch n.direction {
			case 0:
				n.northFerry += ins.value
			case 90:
				n.eastFerry += ins.value
			case 180:
				n.northFerry -= ins.value
			case 270:
				n.eastFerry -= ins.value
			default:
				panic(fmt.Errorf("invalid direction value for ferry %v\n", n.direction))
			}
		}
	}
	if n.eastFerry < 0 {
		n.eastFerry *= -1
	}
	if n.northFerry < 0 {
		n.northFerry *= -1
	}
	return n.northFerry + n.eastFerry
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
