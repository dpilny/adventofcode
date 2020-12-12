package day11

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	seatEmpty = iota + 1
	seatOccupied
	floor
)

type direction struct {
	x int
	y int
}

type seating struct {
	seats           [][]byte
	directNeighbors bool
}

func parse(seatingPath string, directNeighbors bool) (*seating, error) {
	content, err := ioutil.ReadFile(seatingPath)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(content), "\n")
	var data [][]byte
	for _, rawLine := range raw {
		var row []byte
		for _, codePoint := range rawLine {
			switch string(codePoint) {
			case "L":
				row = append(row, seatEmpty)
			case "#":
				row = append(row, seatOccupied)
			case ".":
				row = append(row, floor)
			default:
				panic(fmt.Errorf("unknown seat typs %v", codePoint))
			}
		}
		data = append(data, row)
	}
	return &seating{seats: data, directNeighbors: directNeighbors}, nil
}

func duplicateSeating(seating [][]byte) [][]byte {
	duplicate := make([][]byte, len(seating))
	for i := range seating {
		duplicate[i] = make([]byte, len(seating[i]))
		copy(duplicate[i], seating[i])
	}
	return duplicate
}

func (s *seating) printSeating() {
	seatSymbols := map[byte]string{
		seatEmpty:    "L",
		seatOccupied: "#",
		floor:        ".",
	}
	for _, row := range s.seats {
		for _, seat := range row {
			fmt.Print(seatSymbols[seat])
		}
		fmt.Print("\n")
	}
}

// calculates the seating for people until it has settled
// returns the amount of generation it took until the seating settled and how many seats are occupied
func (s *seating) letPeopleSit() (int, int) {
	gen := 0
	log.Println("letting people sit - initial layout is")
	s.printSeating()
	for {
		gen++
		settled := s.calcNext()
		if settled {
			log.Printf("final seating after gen %v is:\n", gen)
			s.printSeating()
			break
		}
	}
	return gen, s.countOccupied()
}

// calculates the next generation of seats
// returns true if the newly calculated seating has not changed
func (s *seating) calcNext() bool {
	copiedSeats := duplicateSeating(s.seats)
	changed := false
	for rowIndex, row := range copiedSeats {
		for colIndex, seat := range row {
			neighbors := s.getNeighbors(copiedSeats, rowIndex, colIndex)
			newSeat := transformSeat(seat, neighbors, s.directNeighbors)
			if newSeat != seat {
				s.seats[rowIndex][colIndex] = newSeat
				changed = true
			}
		}
	}
	return !changed
}

func (s *seating) getNeighbors(seating [][]byte, row, column int) []byte {
	var neighbors []byte
	if s.directNeighbors {
		// * * *
		// * + *
		// * * *
		if column > 0 {
			neighbors = append(neighbors, seating[row][column-1])
			if row > 0 {
				neighbors = append(neighbors, seating[row-1][column-1])
			}
			if row < len(seating)-1 {
				neighbors = append(neighbors, seating[row+1][column-1])
			}
		}
		if column < len(seating[row])-1 {
			neighbors = append(neighbors, seating[row][column+1])
			if row > 0 {
				neighbors = append(neighbors, seating[row-1][column+1])
			}
			if row < len(seating)-1 {
				neighbors = append(neighbors, seating[row+1][column+1])
			}
		}
		if row > 0 {
			neighbors = append(neighbors, seating[row-1][column])
		}
		if row < len(seating)-1 {
			neighbors = append(neighbors, seating[row+1][column])
		}
	} else {
		/*
		        (-1, -1),
		        (-1, 0),
		        (-1, 1),
		        (0, -1),
		//        (0, 0),
		        (0, 1),
		        (1, -1),
		        (1, 0),
		        (1, 1),
		 */

		// top left
		if seat := getNextSeat(seating, row, column, -1, -1); seat != nil {
			neighbors = append(neighbors, *seat)
		}

		// top
		if seat := getNextSeat(seating, row, column, -1, 0); seat != nil {
			neighbors = append(neighbors, *seat)
		}

		// top right
		if seat := getNextSeat(seating, row, column, -1, 1); seat != nil {
			neighbors = append(neighbors, *seat)
		}

		// left
		if seat := getNextSeat(seating, row, column, 0, -1); seat != nil {
			neighbors = append(neighbors, *seat)
		}

		// right
		if seat := getNextSeat(seating, row, column, 0, 1); seat != nil {
			neighbors = append(neighbors, *seat)
		}

		// bottom left
		if seat := getNextSeat(seating, row, column, 1, -1); seat != nil {
			neighbors = append(neighbors, *seat)
		}

		// bottom
		if seat := getNextSeat(seating, row, column, 1, 0); seat != nil {
			neighbors = append(neighbors, *seat)
		}

		// bottom right
		if seat := getNextSeat(seating, row, column, 1, 1); seat != nil {
			neighbors = append(neighbors, *seat)
		}
	}
	return neighbors
}

func getNextSeat(seating [][]byte, row, column, rowDir, colDir int) *byte {
	r := row + rowDir
	c := column + colDir
	for {
		if r >= len(seating) || r < 0 || c >= len(seating[row]) || c < 0 {
			break
		}
		if seat := seating[r][c]; seat != floor {
			return &seat
		}
		c += colDir
		r += rowDir
	}
	return nil
}

func transformSeat(seat byte, neighbors []byte, directNeighbors bool) byte {
	occupiedSeats := countOccupied(neighbors)
	switch seat {
	case floor:
		return floor
	case seatEmpty:
		if occupiedSeats == 0 {
			return seatOccupied
		} else {
			return seatEmpty
		}
	case seatOccupied:
		if directNeighbors {
			if occupiedSeats >= 4 {
				return seatEmpty
			} else {
				return seatOccupied
			}
		} else {
			if occupiedSeats >= 5 {
				return seatEmpty
			} else {
				return seatOccupied
			}
		}
	default:
		panic(fmt.Errorf("invalid seat type: %v", seat))
	}
}

func (s *seating) countOccupied() int {
	occupied := 0

	for _, row := range s.seats {
		for _, v := range row {
			if v == seatOccupied {
				occupied++
			}
		}
	}

	return occupied
}

func countOccupied(seats []byte) int {
	occupied := 0

	for _, v := range seats {
		if v == seatOccupied {
			occupied++
		}
	}

	return occupied
}
