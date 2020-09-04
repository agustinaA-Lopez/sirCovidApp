package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var jsonText string

type Objetos struct {
	Y      float64
	Height float64
	X      float64
	Width  float64
}

var decoded []Objetos

func main() {
	fJSON, err := ioutil.ReadFile(`sircovid\data\ciudadFereUnDetalle.json`)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(fJSON, &decoded)

	objeto := make([][4]int, len(decoded))

	for i := 0; i < len(decoded); i++ {
		objeto[i][0] = int(decoded[i].X)
		objeto[i][1] = int(decoded[i].Width)
		objeto[i][2] = int(decoded[i].Y)
		objeto[i][3] = int(decoded[i].Height)
	}

	fmt.Println("objetos := make([][]int,", len(objeto), ")")

	for i := 0; i < len(objeto); i++ {
		fmt.Println("objetos[", i, "] = []int{", objeto[i][0], ",", objeto[i][1], ",", objeto[i][2], ",", objeto[i][3], "}")
	}

}
