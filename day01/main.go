package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolvePartOne(intArr []int) int {
	var count int
	for i := 1; i < len(intArr); i++ {
		l := intArr[i-1]
		r := intArr[i]
		if l < r {
			count++
		}
	}

	return count
}

func SolvePartTwo(intArr []int) int {
	var count int
	currSum := intArr[0] + intArr[1] + intArr[2]

	for i := 3; i < len(intArr); i++ {
		newSum := intArr[i-2] + intArr[i-1] + intArr[i]

		if newSum > currSum {
			count++
		}
		currSum = newSum
	}

	return count

}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArr, err := common.ParseIntArray(data, []byte{'\n'})
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(intArr)
	fmt.Println(res)

	res2 := SolvePartTwo(intArr)
	fmt.Println(res2)

}
