package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Cavern struct {
	Grid [][]int
	X    int
	Y    int
}

func (c Cavern) TotalRisk() int {
	var res int
	for x := range c.X {
		for y := range c.Y {
			res += c.Grid[y][x]
		}
	}
	return res
}

func (c *Cavern) Expand(n int) {

	// Expand X
	for i, row := range c.Grid {
		var extension []int
		for j := range n {
			ext := GetIncrementedRow(row, j+1)
			extension = slices.Concat(extension, ext)
		}
		c.Grid[i] = slices.Concat(row, extension)
	}

	i := 0
	YMax := c.Y * n

	for i < YMax {
		newRow := GetIncrementedRow(c.Grid[i], 1)
		c.Grid = append(c.Grid, newRow)
		i++
	}

	c.Y = len(c.Grid)
	c.X = len(c.Grid[0])

}

func GetIncrementedRow(row []int, n int) []int {
	res := make([]int, len(row))
	for i, val := range row {
		div := (val + n) / 10
		newVal := (val+n)%10 + div

		res[i] = newVal
	}

	return res

}

type PosMemo struct {
	models.Pos
	Val int
}

func CreateCavern(data []byte) Cavern {
	var grid [][]int

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		var row []int
		for _, e := range entry {
			row = append(row, int(e-'0'))
		}
		grid = append(grid, row)
	}

	Y := len(grid)
	X := len(grid[0])

	return Cavern{
		Grid: grid,
		X:    X,
		Y:    Y,
	}

}

// Too slow
func SolvePartOneNoMemo(cavern Cavern) int {

	deltas := []models.Pos{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	dst := models.Pos{
		X: cavern.X - 1,
		Y: cavern.Y - 1,
	}

	res := cavern.TotalRisk()

	seen := make(map[models.Pos]bool)

	// DFS without memo
	var dfs func(curr models.Pos, total int)
	dfs = func(curr models.Pos, total int) {

		// exit condition: out of bounds
		if !(0 <= curr.X && curr.X < cavern.X) || !(0 <= curr.Y && curr.Y < cavern.Y) {
			return
		}

		// exit condition: already visited
		if seen[curr] {
			return
		}

		total += cavern.Grid[curr.Y][curr.X]
		// exit condition: hit dst
		if curr == dst {
			res = min(res, total)
			return
		}

		// visit and continue dfs
		seen[curr] = true
		for _, d := range deltas {
			dfs(models.Pos{X: curr.X + d.X, Y: curr.Y + d.Y}, total)
		}
		seen[curr] = false

	}
	dfs(models.Pos{X: 0, Y: 0}, 0)

	return (res - cavern.Grid[0][0])

}

// Similar to LeetCode 64: Miniminum Path Sum
// Only works for part 1
func SolveMemo(cavern Cavern) int {

	deltas := []models.Pos{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	dst := models.Pos{
		X: cavern.X - 1,
		Y: cavern.Y - 1,
	}

	memo := make(map[models.Pos]int)

	var dfs func(curr models.Pos) int
	dfs = func(curr models.Pos) int {

		// case: already solved
		if val, ok := memo[curr]; ok {
			return val
		}

		val := cavern.Grid[curr.Y][curr.X]

		// exit condition: hit dst
		if curr == dst {
			return val
		}

		// visit and continue dfs
		var path *int
		for _, d := range deltas {
			newX := curr.X + d.X
			newY := curr.Y + d.Y
			if (0 <= newX && newX < cavern.X) && (0 <= newY && newY < cavern.Y) {
				pv := dfs(models.Pos{X: newX, Y: newY})
				if path == nil {
					path = &pv
				}
				*path = min(*path, pv)
			}
		}
		memo[curr] = val + *path
		return memo[curr]
	}

	res := dfs(models.Pos{X: 0, Y: 0})

	return res - cavern.Grid[0][0]
}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	c := CreateCavern(data)

	// SolvePartOneNoMemo(c)
	res := SolveMemo(c)
	fmt.Println(res)

	// Part Two
	c.Expand(4)
	res2 := SolveMemo(c)
	fmt.Println(res2)

}
