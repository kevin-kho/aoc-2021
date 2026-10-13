package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

func Insert(original string, newElements []string) string {

	var sb strings.Builder
	i := 0
	j := 0
	k := 0 // even take original odd take newElements

	for k < len(original)+len(newElements) {
		switch k % 2 {
		case 0:
			sb.WriteByte(original[i])
			i++
		case 1:
			sb.WriteString(newElements[j])
			j++
		}
		k++
	}

	return sb.String()

}

func CreateMap(data []byte) map[string]string {
	res := make(map[string]string)
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		src, dst, _ := strings.Cut(string(entry), " -> ")
		res[src] = dst
	}
	return res
}

func SolvePartOne(original string, mp map[string]string) int {

	for range 10 {

		var newElements []string
		for i := 1; i < len(original); i++ {
			pair := original[i-1 : i+1]
			newElements = append(newElements, mp[pair])
		}

		original = Insert(original, newElements)
	}

	ct := make(map[rune]int)
	for _, r := range original {
		ct[r]++
	}
	mn := len(original)
	mx := 0
	for _, n := range ct {
		mn = min(mn, n)
		mx = max(mx, n)
	}

	return mx - mn

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	mp := CreateMap(data)

	// res := SolvePartOne("NNCB", mp)
	res := SolvePartOne("PBVHVOCOCFFNBCNCCBHK", mp)

	fmt.Println(res)
}
