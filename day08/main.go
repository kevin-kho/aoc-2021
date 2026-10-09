package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Digit struct {
	Val       int
	StringRep string
	Runes     map[rune]bool
}

func CreateDigit(val int, chars string) Digit {

	runes := make(map[rune]bool)

	for _, c := range chars {
		runes[c] = true
	}

	return Digit{
		Val:       val,
		StringRep: chars,
		Runes:     runes,
	}

}

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

func SolvePartTwo(signals []Signal) {
	var patterns []string
	var outputs []string
	digitMap := make(map[int]Digit)

	// for _, s := range signals {
	// 	patterns = slices.Concat(patterns, s.Pattern)
	// 	outputs = slices.Concat(outputs, s.Output)
	// }

	patterns = append(patterns, signals[0].Pattern...)
	outputs = append(outputs, signals[0].Output...)

	combined := slices.Concat(patterns, outputs)

	mp := make(map[int]map[string]bool) // key: len, value: pattern & output

	for _, c := range combined {
		if mp[len(c)] == nil {
			mp[len(c)] = make(map[string]bool)

		}
		r := []rune(c)
		slices.Sort(r)
		c = string(r)
		mp[len(c)][c] = true
	}

	// Assign unique ones: 1, 4, 7 , 8
	for length, digit := range GetUniqueDigitMap() {
		for k := range mp[length] {
			digitMap[digit] = CreateDigit(digit, k)
		}
	}

	for length, vals := range mp {
		fmt.Println(length, vals)
	}

	// Determine 3
	for str := range mp[5] {
		found := true
		for r := range digitMap[1].Runes {
			if !strings.ContainsRune(str, r) {
				found = false
			}
		}
		if found {
			digitMap[3] = CreateDigit(3, str)
			delete(mp[5], str)
			break
		}
	}

	// Determine 5
	// It will have all chars from 4 except one which is missing from 1 as well
	for str := range mp[5] {
		var missing []rune
		for r := range digitMap[4].Runes {
			if !strings.ContainsRune(str, r) {
				missing = append(missing, r)
			}
		}

		switch len(missing) {
		case 1:
			digitMap[5] = CreateDigit(5, str)
		case 2:
			digitMap[2] = CreateDigit(2, str)
		}

	}

	for val, digit := range digitMap {
		fmt.Println(val, digit)
	}

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("input.txt")
	data, err := common.ReadInput("inputExamplePartTwo.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	signals := GetSignal(data)

	res := SolvePartOne(signals)
	fmt.Println(res)

	SolvePartTwo(signals)

}
