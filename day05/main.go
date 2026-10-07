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

func (l Line) GetVector() Pos {
	vecX := l.End.X - l.Start.X
	vecY := l.End.Y - l.Start.Y

	return Pos{
		X: vecX,
		Y: vecY,
	}

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

func FilterStraightLines(lines []Line) []Line {
	var res []Line
	for _, l := range lines {
		st := l.Start
		ed := l.End

		if st.X == ed.X || st.Y == ed.Y {
			res = append(res, l)
		}
	}

	return res

}

func SolvePartOne(lines []Line) int {

	mp := make(map[Pos]int)

	straightLines := FilterStraightLines(lines)
	for _, l := range straightLines {

		vec := l.GetVector()
		// either dx or dy is guaranteed to be 0
		dx := common.IntAbs(vec.X)
		dy := common.IntAbs(vec.Y)

		p := Pos{
			X: min(l.Start.X, l.End.X),
			Y: min(l.Start.Y, l.End.Y),
		}

		mp[p]++

		for range dx {
			p.X += 1
			mp[p]++
		}

		for range dy {
			p.Y += 1
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

}
