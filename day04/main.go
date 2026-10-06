package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

type Board struct {
	ValMp   map[int]Pos // if it's used, delete from map
	RowMp   map[int]int // key: row#, value: count (start at 5)
	ColMp   map[int]int // key: col#, value: count (start at 5)
	UsedVal []int
}

func (b Board) TotalizeUnusedNumbers() int {
	var res int

	for val := range b.ValMp {
		res += val
	}

	return res

}

func GetNumbers(data []byte) ([]int, error) {
	var res []int
	for entryStr := range strings.SplitSeq(string(data), ",") {

		entryInt, err := strconv.Atoi(entryStr)
		if err != nil {
			return res, err
		}
		res = append(res, entryInt)
	}

	return res, nil

}

func CreateBoard(board []string) (Board, error) {
	valMp := make(map[int]Pos)
	for y, row := range board {
		row = strings.TrimSpace(row)
		var formattedRow []string
		for entry := range strings.SplitSeq(row, " ") {
			if entry == "" {
				continue
			}
			formattedRow = append(formattedRow, entry)
		}
		for x, val := range formattedRow {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				return Board{}, err
			}

			valMp[valInt] = Pos{
				X: x,
				Y: y,
			}
		}
	}

	rowMp := make(map[int]int)
	colMp := make(map[int]int)
	for i := range 5 {
		rowMp[i] = 5
		colMp[i] = 5
	}

	return Board{
		ValMp: valMp,
		RowMp: rowMp,
		ColMp: colMp,
	}, nil

}

func GetBoards(data []byte) ([]Board, error) {
	var res []Board

	var boards [][]string
	var curr []string
	for entry := range strings.SplitSeq(string(data), "\n") {
		// fmt.Println(entry)
		if entry == "" {
			boards = append(boards, curr)
			curr = []string{}
			continue
		}
		curr = append(curr, entry)
	}
	if len(curr) > 0 {
		boards = append(boards, curr)
	}

	for _, board := range boards {
		brd, err := CreateBoard(board)
		if err != nil {
			return res, err
		}
		res = append(res, brd)
	}

	return res, nil

}

func SolvePartOne(numbers []int, boards []Board) int {

	for _, num := range numbers {
		for _, board := range boards {

			if pos, ok := board.ValMp[num]; ok {
				delete(board.ValMp, num)

				if board.RowMp[pos.Y] == 1 || board.ColMp[pos.X] == 1 {
					return num * board.TotalizeUnusedNumbers()
				}

				board.RowMp[pos.Y]--
				board.ColMp[pos.X]--
			}

		}

	}

	// No board won
	return -1

}

func main() {
	// numberData, err := common.ReadInput("inputNumbersExample.txt")
	numberData, err := common.ReadInput("inputNumbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	numberData = common.TrimNewLineSuffix(numberData)
	numbers, err := GetNumbers(numberData)
	if err != nil {
		log.Fatal(err)
	}

	// boardData, err := common.ReadInput("inputBoardsExample.txt")
	boardData, err := common.ReadInput("inputBoards.txt")
	if err != nil {
		log.Fatal(err)
	}
	boardData = common.TrimNewLineSuffix(boardData)
	boards, err := GetBoards(boardData)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(numbers, boards)
	fmt.Println(res)

}
