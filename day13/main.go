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

}

func (p *Paper) FoldAlongY(y int) {

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

func main() {

	pointData, err := common.ReadInput("inputPointsExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	pointData = common.TrimNewLineSuffix(pointData)

	foldData, err := common.ReadInput("inputFoldsExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	foldData = common.TrimNewLineSuffix(foldData)

	paper, err := CreatePaper(pointData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(paper)

	folds, err := CreateFolds(foldData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folds)

}
