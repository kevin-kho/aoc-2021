package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Line struct {
	Start Pos
	End   Pos
}

type Pos struct {
	X int
	Y int
}

func CreatePos(posStr string) (Pos, error) {
	var res Pos
	xStr, yStr, _ := strings.Cut(posStr, ",")
	xInt, err := strconv.Atoi(xStr)
	if err != nil {
		return res, nil
	}

	yInt, err := strconv.Atoi(yStr)
	if err != nil {
		return res, nil
	}

	res.X = xInt
	res.Y = yInt
	return res, nil

}

func CreateLines(data []byte) ([]Line, error) {
	var res []Line
	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		st, ed, _ := strings.Cut(string(entry), " -> ")

		start, err := CreatePos(st)
		if err != nil {
			return res, err
		}

		end, err := CreatePos(ed)
		if err != nil {
			return res, err
		}

		res = append(res, Line{
			Start: start,
			End:   end,
		})

	}

	return res, nil

}

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	lines, err := CreateLines(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)

}
