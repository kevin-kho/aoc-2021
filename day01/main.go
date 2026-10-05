package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetIntArr(data []byte) ([]int, error) {
	var res []int

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)
		i, err := strconv.Atoi(string(entry))
		if err != nil {
			return res, err
		}

		res = append(res, i)

	}

	return res, nil

}

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

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArr, err := GetIntArr(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(intArr)
	fmt.Println(res)

}
