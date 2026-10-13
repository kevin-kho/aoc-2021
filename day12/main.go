package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/kevin-kho/aoc-utilities/common"
)

func BuildAdjList(data []byte) map[string][]string {
	res := make(map[string][]string)
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		src, dst, _ := strings.Cut(string(entry), "-")
		res[src] = append(res[src], dst)
		res[dst] = append(res[dst], src)
	}

	return res

}

func SolvePartOne(adj map[string][]string) {

	seen := make(map[string]bool)

	var dfs func(curr string) int
	dfs = func(curr string) int {

		// exit condition: reach the end
		if curr == "end" {
			return 1
		}

		// exit condition: already seen
		if seen[curr] {
			return 0
		}

		// set visited if small cave
		if unicode.IsLower(rune(curr[0])) {
			seen[curr] = true
		}
		var val int
		for _, nei := range adj[curr] {
			val += dfs(nei)
		}
		seen[curr] = false

		return val

	}

	res := dfs("start")
	fmt.Println(res)

}

func SolvePartTwo(adj map[string][]string) {

	seen := make(map[string]bool)

	var dfs func(curr string, smallCaveVisited string) int
	dfs = func(curr string, smallCaveVisited string) int {

		// exit condition: reach the end
		if curr == "end" {
			return 1
		}

		if curr == "start" && seen[curr] {
			return 0
		}

		// exit condition
		if seen[curr] && smallCaveVisited != "" {
			return 0
		}

		if seen[curr] {
			smallCaveVisited = curr
		}

		// set visited if small cave
		if unicode.IsLower(rune(curr[0])) {
			seen[curr] = true
		}
		var val int
		for _, nei := range adj[curr] {
			val += dfs(nei, smallCaveVisited)
		}

		if smallCaveVisited != curr {
			seen[curr] = false
		}

		return val
	}

	res := dfs("start", "")
	fmt.Println(res)

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	// data, err := common.ReadInput("inputExample2.txt")
	// data, err := common.ReadInput("inputExample3.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	adj := BuildAdjList(data)

	SolvePartOne(adj)

	SolvePartTwo(adj)

}
