package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
	"github.com/kevin-kho/aoc-utilities/models"
)

type Octopus struct {
	models.Pos
	Energy int
}

func GetOctopus(data []byte) {
	res := make(map[models.Pos]Octopus)
	for y, entry := range bytes.Split(data, []byte{'\n'}) {
		for x, b := range entry {

			p := models.Pos{
				X: x,
				Y: y,
			}

			res[p] = Octopus{
				Pos:    p,
				Energy: int(b - '0'),
			}

		}
	}

	fmt.Println(res)

}

func main() {

	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	GetOctopus(data)

}
