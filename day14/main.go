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

func Solve(original string, mp map[string]string, steps int) int {

	for range steps {

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

func SolveDfs(original string, mp map[string]string, steps int) int {

	count := make(map[string]int)
	for _, char := range original {
		count[string(char)]++
	}

	var dfs func(l, r string, step int)
	dfs = func(l, r string, step int) {
		if step == steps {
			return
		}

		mid := mp[l+r]
		count[mid]++

		dfs(l, mid, step+1)
		dfs(mid, r, step+1)

	}

	for i := 1; i < len(original); i++ {
		l := string(original[i-1])
		r := string(original[i])
		dfs(l, r, 0)
	}

	// set mn and mx
	var mn int
	var mx int
	for _, n := range count {
		mn = n
		mx = n
		break
	}

	// find it
	for _, n := range count {
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

	// res := Solve("NNCB", mp, 10)
	res := Solve("PBVHVOCOCFFNBCNCCBHK", mp, 10)
	fmt.Println(res)

	// resDfs := SolveDfs("NNCB", mp, 10)
	resDfs := SolveDfs("PBVHVOCOCFFNBCNCCBHK", mp, 10)
	fmt.Println(resDfs)

}
