package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Fish struct {
	Timer int
	Day   int
}

func Solve(intArr []int, day int) int {

	var fish int

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
	return fish

}

func SolveMemo(intArr []int, day int) int {

	var fish int

	mp := make(map[Fish]int)

	var dfs func(fish Fish) int
	dfs = func(fish Fish) int {
		if fish.Day == day {
			return 1
		}

		if val, ok := mp[fish]; ok {
			return val
		}

		var res int
		if fish.Timer == 0 {
			res += dfs(Fish{
				Timer: 6,
				Day:   fish.Day + 1,
			})
			res += dfs(Fish{
				Timer: 8,
				Day:   fish.Day + 1,
			})
		} else {
			res += dfs(Fish{
				Timer: fish.Timer - 1,
				Day:   fish.Day + 1,
			})
		}

		mp[fish] = res
		return mp[fish]
	}
	for _, i := range intArr {
		fish += dfs(Fish{
			Timer: i,
			Day:   0,
		})
	}

	return fish

}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArr, err := common.ParseIntArray(data, []byte{','})
	if err != nil {
		log.Fatal(err)
	}

	res := Solve(intArr, 80)
	fmt.Println(res)

	resMemo := SolveMemo(intArr, 80)
	fmt.Println(resMemo)

	res2 := SolveMemo(intArr, 256)
	fmt.Println(res2)

}
