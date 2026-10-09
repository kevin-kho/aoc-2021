package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Grid struct {
	X     int
	Y     int
	Board [][]byte
}

func (g Grid) GetLowPoints() []LowPoint {
	deltas := models.GetDeltas()
	var res []LowPoint

	for x := range g.X {
		for y := range g.Y {

			val := g.Board[y][x]
			lowPoint := true
			for _, d := range deltas {
				newX, newY := x+d.X, y+d.Y
				if !(0 <= newX && newX < g.X) || !(0 <= newY && newY < g.Y) {
					continue
				}
				if val >= g.Board[newY][newX] {
					lowPoint = false
				}

			}
			if lowPoint {
				res = append(res, LowPoint{
					X:   x,
					Y:   y,
					Val: int(val - '0'),
				})
			}

		}
	}

	return res
}

type LowPoint struct {
	models.Pos
	Val int
}

func GetGrid(data []byte) Grid {
	board := bytes.Split(data, []byte{'\n'})
	x := len(board[0])
	y := len(board)

	return Grid{
		X:     x,
		Y:     y,
		Board: board,
	}
}

func SolvePartOne(grid Grid) int {

	lowPoints := grid.GetLowPoints()

	var res int
	for _, lp := range lowPoints {
		res += lp.Val + 1
	}

	return res

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	grid := GetGrid(data)
	res := SolvePartOne(grid)
	fmt.Println(res)

}
