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

type Line struct {
	Start models.Pos
	End   models.Pos
}

func (l Line) GetVector() models.Pos {
	vecX := l.End.X - l.Start.X
	vecY := l.End.Y - l.Start.Y

	return models.Pos{
		X: vecX,
		Y: vecY,
	}

}

func CreatePos(posStr string) (models.Pos, error) {
	var res models.Pos
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

func FilterLines(lines []Line) ([]Line, []Line) {
	var straight []Line
	var diagonal []Line
	for _, l := range lines {
		st := l.Start
		ed := l.End

		if st.X == ed.X || st.Y == ed.Y {
			straight = append(straight, l)
		} else {
			diagonal = append(diagonal, l)
		}
	}

	return straight, diagonal

}

func SolvePartOne(lines []Line) int {

	mp := make(map[models.Pos]int)

	straightLines, _ := FilterLines(lines)
	for _, l := range straightLines {

		d := l.GetVector()
		gcd := d.GetGcd()

		d.X = d.X / gcd
		d.Y = d.Y / gcd

		p := l.Start

		mp[p]++

		for p != l.End {
			p.X += d.X
			p.Y += d.Y

			mp[p]++
		}

	}

	var count int
	for _, ct := range mp {
		if ct > 1 {
			count++
		}
	}

	return count

}

func SolvePartTwo(lines []Line) int {
	mp := make(map[models.Pos]int)

	for _, l := range lines {
		d := l.GetVector()
		gcd := d.GetGcd()

		d.X = d.X / gcd
		d.Y = d.Y / gcd

		p := l.Start

		mp[p]++
		for p != l.End {
			p.X += d.X
			p.Y += d.Y
			mp[p]++
		}
	}

	var count int
	for _, ct := range mp {
		if ct > 1 {
			count++
		}
	}

	return count

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	lines, err := CreateLines(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(lines)
	fmt.Println(res)

	res2 := SolvePartTwo(lines)
	fmt.Println(res2)

}
