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

func GetAtmosphere(binIntArr BinIntArr) int64 {

	oxy := slices.Clone(binIntArr.Arr)
	co2 := slices.Clone(binIntArr.Arr)

	for i := range binIntArr.BitSize {
		if len(oxy) == 1 {
			break
		}

		// handle oxygen
		oxyMp := make(map[byte][]BinInt)
		for _, num := range oxy {
			char := num.Str[i]
			oxyMp[char] = append(oxyMp[char], num)
		}
		// determine next oxygen
		switch len(oxyMp['1']) >= len(oxyMp['0']) {
		case true:
			oxy = oxyMp['1']
		case false:
			oxy = oxyMp['0']
		}
	}

	for i := range binIntArr.BitSize {
		if len(co2) == 1 {
			break
		}

		// handle co2
		co2Mp := make(map[byte][]BinInt)
		for _, num := range co2 {
			char := num.Str[i]
			co2Mp[char] = append(co2Mp[char], num)
		}

		// determine next co2
		switch len(co2Mp['0']) <= len(co2Mp['1']) {
		case true:
			co2 = co2Mp['0']
		case false:
			co2 = co2Mp['1']
		}

	}

	return oxy[0].Int * co2[0].Int

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

	res2 := GetAtmosphere(binIntArr)
	fmt.Println(res2)

}
