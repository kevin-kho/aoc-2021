package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetIntArr(data []byte) ([]int, error) {
	var res []int
	for entry := range strings.SplitSeq(string(data), ",") {
		i, err := strconv.Atoi(entry)
		if err != nil {
			return res, err
		}

		res = append(res, i)
	}

	return res, nil
}

func SolvePartOne(intArr []int) {

	var fish int
	day := 80

	var dfs func(timer int, d int)
	dfs = func(timer int, d int) {
		if d == day {
			fish++
			return
		}

		if timer == 0 {
			dfs(6, d+1) // reset timer
			dfs(8, d+1) // new fish
		} else {
			dfs(timer-1, d+1) // countdown timer
		}
	}
	for _, i := range intArr {
		dfs(i, 0)
	}
	fmt.Println(fish)

}

func main() {

	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArr, err := GetIntArr(data)
	if err != nil {
		log.Fatal(err)
	}

	SolvePartOne(intArr)

}
