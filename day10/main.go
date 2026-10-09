package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Chunk struct {
	Chars            []byte
	Illegal          bool
	FirstIllegalChar byte
}

func CreateChunk(entry []byte) Chunk {
	opener := map[byte]bool{
		'(': true,
		'[': true,
		'{': true,
		'<': true,
	}
	closer := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	var stack []byte
	for _, b := range entry {
		// case: opener
		if opener[b] {
			stack = append(stack, b)
			continue
		}

		// case: closer and empty stack
		if len(stack) == 0 {
			return Chunk{
				Chars:            entry,
				Illegal:          true,
				FirstIllegalChar: b,
			}
		}

		// case: closer & illegal
		if closer[b] != stack[len(stack)-1] {
			return Chunk{
				Chars:            entry,
				Illegal:          true,
				FirstIllegalChar: b,
			}
		}

		stack = stack[:len(stack)-1]

	}

	return Chunk{
		Chars:            entry,
		Illegal:          false,
		FirstIllegalChar: 0,
	}

}

func GetChunks(data []byte) []Chunk {
	var res []Chunk

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		res = append(res, CreateChunk(entry))
	}

	return res

}

func SolvePartOne(chunks []Chunk) int {
	var res int
	score := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, c := range chunks {
		if c.Illegal {
			res += score[c.FirstIllegalChar]
		}
	}

	return res
}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	chunks := GetChunks(data)

	res := SolvePartOne(chunks)
	fmt.Println(res)

}
