package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Board struct {
	Octopi map[models.Pos]*Octopus
	X      int
	Y      int
}

type Octopus struct {
	models.Pos
	Energy int
}

func (o *Octopus) IncreaseEnergy() {
	o.Energy += 1
	o.Energy = min(10, o.Energy)
}

func (o *Octopus) ResetEnergy() {
	o.Energy = 0
}

func PrintBoard(octopi map[models.Pos]*Octopus) {
	var X int
	var Y int
	for p := range octopi {
		X = max(X, p.X)
		Y = max(Y, p.Y)
	}

	X++
	Y++

	for y := range Y {
		var row []int
		for x := range X {
			row = append(row, octopi[models.Pos{X: x, Y: y}].Energy)
		}
		fmt.Println(row)
	}

}

func GetBoard(data []byte) Board {
	res := make(map[models.Pos]*Octopus)
	var X int
	var Y int
	for y, entry := range bytes.Split(data, []byte{'\n'}) {
		for x, b := range entry {

			p := models.Pos{
				X: x,
				Y: y,
			}

			res[p] = &Octopus{
				Pos:    p,
				Energy: int(b - '0'),
			}
			X = max(X, x)
		}
		Y = max(Y, y)
	}

	return Board{
		Octopi: res,
		X:      X,
		Y:      Y,
	}

}

func SolvePartOne(board Board) {
	deltas := models.GetDeltas()
	deltas = append(deltas,
		models.Pos{X: -1, Y: 1},
		models.Pos{X: -1, Y: -1},
		models.Pos{X: 1, Y: 1},
		models.Pos{X: 1, Y: -1},
	)

	for range 1 {
		mp := make(map[models.Pos]int) // count of how many times to increase energy
		for p, o := range board.Octopi {
			mp[p]++
			if o.Energy == 9 {
				for _, d := range deltas {
					pNei := models.Pos{
						X: p.X + d.X,
						Y: p.Y + d.Y}
					if !(0 <= pNei.X && pNei.X < board.X) || !(0 <= pNei.Y && pNei.Y < board.Y) {
						continue
					}
					mp[pNei]++
				}
			}
		}

	}
	PrintBoard(board.Octopi)

}

func main() {

	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	board := GetBoard(data)

	SolvePartOne(board)

}
