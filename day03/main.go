package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

type BinInt struct {
	Str string
	Int int64
}

type BinIntArr struct {
	Arr     []BinInt
	BitSize int
}

type GammaEpsilon struct {
	Gamma   int64
	Epsilon int64
}

type Atmosphere struct {
	Oxygen int64
	CO2    int64
}

func GetBinIntArr(data []byte) (BinIntArr, error) {
	var res []BinInt
	var bitSize int

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)
		bitSize = len(string(entry))
		i, err := strconv.ParseInt(string(entry), 2, 16)
		if err != nil {
			return BinIntArr{}, err
		}

		res = append(res, BinInt{
			Str: string(entry),
			Int: i,
		})
	}

	return BinIntArr{
		Arr:     res,
		BitSize: bitSize,
	}, nil

}

func GetGammaEpsilon(binIntArr BinIntArr) (GammaEpsilon, error) {

	n := len(binIntArr.Arr)
	binIntArr.Arr = slices.Clone(binIntArr.Arr)

	var gammaByte []byte
	var epsilonByte []byte

	for range binIntArr.BitSize {
		var count int64

		for i, binInt := range binIntArr.Arr {
			count += binInt.Int & 1
			binInt.Int = binInt.Int >> 1
			binIntArr.Arr[i] = binInt
		}

		if count > int64(n)/2 {
			gammaByte = append(gammaByte, '1')
			epsilonByte = append(epsilonByte, '0')
		} else {
			gammaByte = append(gammaByte, '0')
			epsilonByte = append(epsilonByte, '1')
		}

	}

	slices.Reverse(gammaByte)
	slices.Reverse(epsilonByte)

	gamma, err := strconv.ParseInt(string(gammaByte), 2, 16)
	if err != nil {
		return GammaEpsilon{}, err
	}

	epsilon, err := strconv.ParseInt(string(epsilonByte), 2, 16)
	if err != nil {
		return GammaEpsilon{}, err
	}

	return GammaEpsilon{
		Gamma:   gamma,
		Epsilon: epsilon,
	}, nil

}

func GetAtmosphere(binIntArr BinIntArr) {

	binIntArr.Arr = slices.Clone(binIntArr.Arr)
}

func SolvePartOne(binIntArr BinIntArr) (int64, error) {
	var res int64
	ge, err := GetGammaEpsilon(binIntArr)
	if err != nil {
		return res, err
	}

	return ge.Gamma * ge.Epsilon, nil

}

func SolvePartTwo(binIntArr BinIntArr) {

}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	binIntArr, err := GetBinIntArr(data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := SolvePartOne(binIntArr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
