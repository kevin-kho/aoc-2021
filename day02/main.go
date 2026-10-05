package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Action string

const (
	ActionForward Action = "forward"
	ActionDown    Action = "down"
	ActionUp      Action = "up"
)

type Pos struct {
	X   int
	Y   int
	Aim int
}

type Command struct {
	Action Action
	Val    int
}

func CreateCommands(data []byte) ([]Command, error) {
	var res []Command

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)
		action, val, _ := strings.Cut(string(entry), " ")

		valInt, err := strconv.Atoi(val)
		if err != nil {
			return res, err
		}

		var act Action
		switch action {
		case "forward":
			act = ActionForward
		case "down":
			act = ActionDown
		case "up":
			act = ActionUp
		}

		res = append(res, Command{
			Action: act,
			Val:    valInt,
		})

	}

	return res, nil

}

func SolvePartOne(cmds []Command) int {
	var pos Pos

	for _, c := range cmds {
		switch c.Action {
		case ActionForward:
			pos.X += c.Val
		case ActionDown:
			pos.Y += c.Val
		case ActionUp:
			pos.Y -= c.Val
		}
	}

	return pos.X * pos.Y

}

func SolvePartTwo(cmds []Command) int {
	var pos Pos

	for _, c := range cmds {
		switch c.Action {
		case ActionForward:
			pos.X += c.Val
			pos.Y += (pos.Aim * c.Val)
		case ActionDown:
			pos.Aim += c.Val
		case ActionUp:
			pos.Aim -= c.Val
		}
	}

	return pos.X * pos.Y

}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	cmds, err := CreateCommands(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(cmds)
	fmt.Println(res)

	res2 := SolvePartTwo(cmds)
	fmt.Println(res2)

}
