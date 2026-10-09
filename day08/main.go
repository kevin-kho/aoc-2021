package main

import (
	"fmt"
	"log"
	"maps"
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

	// Determine 3
	// 3 must contain all of 1
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

	// Determine 5 and 2
	// It will have all chars from 4 except one which is missing from 1 as well
	// 2 will have all except two which are missing
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

	// Determine 9: digit[5] + digit[1] == 9
	nineRunes := maps.Clone(digitMap[5].Runes)
	maps.Copy(nineRunes, digitMap[1].Runes)

	digitMap[9] = Digit{
		Val:       9,
		StringRep: string(slices.Sorted(maps.Keys(nineRunes))),
		Runes:     nineRunes,
	}
	delete(mp[6], digitMap[9].StringRep)

	// Determine 6 and 0
	// 6 fully contains 5
	// remaining one is zero

	for str := range mp[6] {
		containsAll := true
		for r := range digitMap[5].Runes {
			if !strings.ContainsRune(str, r) {
				containsAll = false
			}
		}

		switch containsAll {
		case true:
			digitMap[6] = CreateDigit(6, str)
		case false:
			digitMap[0] = CreateDigit(0, str)

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
