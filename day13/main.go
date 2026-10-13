package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Fold struct {
	Axis string
	Val  int
}

type Paper struct {
	Points map[models.Pos]bool
}

func (p *Paper) FoldAlongX(x int) {
	// Keep everything to the left of X
	// Points are guarantted to never be along X
	keep := make(map[models.Pos]bool)
	toMirror := make(map[models.Pos]bool)
	for pos := range p.Points {
		if pos.X < x {
			keep[pos] = true
		} else {
			toMirror[pos] = true
		}
	}

	// points to the right of x must be mirrored to the left
	// mirror by pos.X -= abs(pos.X - x) * 2
	for pos := range toMirror {
		pos.X -= (common.IntAbs(pos.X-x) * 2)
		keep[pos] = true
	}

	p.Points = keep

}

func (p *Paper) FoldAlongY(y int) {
	// Keep everything above of Y
	// Points are guaranteed to never be along Y

	keep := make(map[models.Pos]bool)
	toMirror := make(map[models.Pos]bool)
	for pos := range p.Points {
		if pos.Y < y {
			keep[pos] = true
		} else {
			toMirror[pos] = true
		}
	}

	// points to the bottom of y must be mirrored up
	// mirror by pos.Y -= abs(pos.Y - y) * 2
	for pos := range toMirror {
		pos.Y -= (common.IntAbs(pos.Y-y) * 2)
		keep[pos] = true
	}

	p.Points = keep

}

func (p Paper) PrintPaper() {
	X := 0
	Y := 0
	for pos := range p.Points {
		X = max(X, pos.X)
		Y = max(Y, pos.Y)
	}
	X++
	Y++

	var grid [][]string
	for range Y {
		row := make([]string, X)
		grid = append(grid, row)
	}

	for y := range Y {
		for x := range X {
			var b string
			if p.Points[models.Pos{X: x, Y: y}] {
				b = "#"
			} else {
				b = "."
			}
			grid[y][x] = b

		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

}

func CreatePoints(data []byte) (map[models.Pos]bool, error) {
	res := make(map[models.Pos]bool)
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		x, y, _ := strings.Cut(string(entry), ",")
		xInt, err := strconv.Atoi(x)
		if err != nil {
			return res, err
		}

		yInt, err := strconv.Atoi(y)
		if err != nil {
			return res, err
		}

		res[models.Pos{
			X: xInt,
			Y: yInt,
		}] = true

	}

	return res, nil

}

func CreatePaper(data []byte) (Paper, error) {

	pt, err := CreatePoints(data)
	if err != nil {
		return Paper{}, err
	}

	return Paper{
		Points: pt,
	}, nil

}

func CreateFolds(data []byte) ([]Fold, error) {

	var res []Fold

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entry = bytes.TrimPrefix(entry, []byte("fold along "))

		axis, val, _ := strings.Cut(string(entry), "=")
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return res, err
		}

		res = append(res, Fold{
			Axis: axis,
			Val:  valInt,
		})
	}

	return res, nil
}

func Solve(paper Paper, folds []Fold) {
	var dotCount []int
	for _, f := range folds {
		switch f.Axis {
		case "x":
			paper.FoldAlongX(f.Val)
		case "y":
			paper.FoldAlongY(f.Val)
		}

		dotCount = append(dotCount, len(paper.Points))
	}

	fmt.Println(dotCount[0]) // Part One

	paper.PrintPaper() // Part Two

}

func main() {

	// pointData, err := common.ReadInput("inputPointsExample.txt")
	pointData, err := common.ReadInput("inputPoints.txt")
	if err != nil {
		log.Fatal(err)
	}
	pointData = common.TrimNewLineSuffix(pointData)

	// foldData, err := common.ReadInput("inputFoldsExample.txt")
	foldData, err := common.ReadInput("inputFolds.txt")
	if err != nil {
		log.Fatal(err)
	}
	foldData = common.TrimNewLineSuffix(foldData)

	paper, err := CreatePaper(pointData)
	if err != nil {
		log.Fatal(err)
	}

	folds, err := CreateFolds(foldData)
	if err != nil {
		log.Fatal(err)
	}

	Solve(paper, folds)

}
