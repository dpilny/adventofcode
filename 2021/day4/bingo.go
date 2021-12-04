package day4

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BingoGame struct {
	drawnNumbers []int
	boards       *[]Board
}

type Board struct {
	initialBoard [][]int
	workingBoard [][]int
}

type SelectedIndex struct {
	x, y int
}

func GetGame(path string) (*BingoGame, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(data), "\n")
	var numbers []int
	var boards []Board

	var currentBoard [][]int
	for idx, rawLine := range raw {
		if idx == 0 {
			numbers, err = convertStringToIntArray(strings.Split(rawLine, ","))
			if err != nil {
				return nil, err
			}
			continue
		}
		if rawLine == "" {
			continue
		}
		boardNumbers, err := convertStringToIntArray(strings.Fields(rawLine))
		if err != nil {
			return nil, err
		}
		currentBoard = append(currentBoard, boardNumbers)
		if len(currentBoard) == 5 {
			workingBoard := make([][]int, len(currentBoard))
			initialBoard := make([][]int, len(currentBoard))
			copyData(currentBoard, workingBoard)
			copyData(currentBoard, initialBoard)
			boards = append(boards, Board{
				initialBoard: initialBoard,
				workingBoard: workingBoard,
			})
			currentBoard = currentBoard[:0]
		}
	}

	return &BingoGame{
		drawnNumbers: numbers,
		boards:       &boards,
	}, nil
}

func copyData(src [][]int, dst [][]int) {
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
}

func convertStringToIntArray(values []string) ([]int, error) {
	var converted []int
	for _, val := range values {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		converted = append(converted, parsed)
	}
	return converted, nil
}

func (g *BingoGame) firstWin() int {
	for round, number := range g.drawnNumbers {
		for boardId, board := range *g.boards {
			won := board.addDrawnNumber(number)
			if won {
				fmt.Printf("board %v won in round %v\n", boardId+1, round+1)
				board.printBoard()
				return board.remainingSum() * number
			}
		}
	}
	return -1
}

func (g *BingoGame) lastWin() int {
	remainingBoards := len(*g.boards)
	wonBoards := map[int]bool{}
	for round, number := range g.drawnNumbers {
		for boardId, board := range *g.boards {
			won := board.addDrawnNumber(number)
			if won && !wonBoards[boardId] {
				wonBoards[boardId] = true
				fmt.Printf("board %v won in round %v with number %v\n", boardId+1, round+1, number)
				board.printBoard()
				remainingBoards--
				if remainingBoards == 0 {
					return board.remainingSum() * number
				}
			}

		}
	}
	return -1
}

func (b *Board) printBoard() {
	fmt.Println("initial board")
	for _, row := range b.initialBoard {
		fmt.Println(row)
	}

	fmt.Println("seleted board")
	for _, row := range b.workingBoard {
		fmt.Println(row)
	}

}

func (b *Board) remainingSum() int {
	sum := 0
	for _, rows := range b.workingBoard {
		for _, value := range rows {
			if value != -1 {
				sum += value
			}
		}
	}
	return sum
}

func (b *Board) addDrawnNumber(number int) bool {
	idx, err := b.isInBoard(number)
	if err == nil {
		b.workingBoard[idx.y][idx.x] = -1
	}
	return b.hasBingo()
}

func (b *Board) hasBingo() bool {
	colSums := map[int]int{}
	rowSums := map[int]int{}
	for y, rows := range b.workingBoard {
		for x, value := range rows {
			if y == 0 {
				colSums[x] = value
			} else {
				colSums[x] = colSums[x] + value
			}
			if x == 0 {
				rowSums[y] = value
			} else {
				rowSums[y] = rowSums[y] + value
			}
		}
	}
	for _, sum := range colSums {
		if sum == -5 {
			return true
		}
	}
	for _, sum := range rowSums {
		if sum == -5 {
			return true
		}
	}
	return false
}

func (b *Board) isInBoard(number int) (*SelectedIndex, error) {
	for y, rows := range b.initialBoard {
		for x, value := range rows {
			if value == number {
				return &SelectedIndex{
					x: x,
					y: y,
				}, nil
			}
		}
	}
	return nil, errors.New("not in initialBoard")
}
