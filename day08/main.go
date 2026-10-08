package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Signal struct {
	Pattern []string
	Output  []string
}

func GetUniqueDigitMap() map[int]int {
	// 1, 4, 7, and 8 use unique number of segments
	// key: # of segment, val: number it represents
	mp := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}

	return mp

}

func GetSignal(data []byte) []Signal {
	var res []Signal

	for entry := range strings.SplitSeq(string(data), "\n") {
		entry = strings.TrimSpace(entry)
		ptrn, out, _ := strings.Cut(entry, " | ")

		ptrnArr := strings.Split(ptrn, " ")
		outArr := strings.Split(out, " ")

		res = append(res, Signal{
			Pattern: ptrnArr,
			Output:  outArr,
		})
	}

	return res
}

func SolvePartOne(signals []Signal) int {
	var outputs []string
	for _, s := range signals {
		outputs = slices.Concat(outputs, s.Output)
	}

	mp := GetUniqueDigitMap()
	var count int
	for _, o := range outputs {
		if _, ok := mp[len(o)]; ok {
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

	signals := GetSignal(data)

	res := SolvePartOne(signals)
	fmt.Println(res)

}
