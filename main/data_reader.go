package main

import (
	"fmt"
	"os"

	"gonum.org/v1/gonum/mat"
)

type data struct {
	raw_data []mat.Dense
	input    []mat.Dense
	labels   []mat.Dense
	name     string
}

func load_data(filepath string) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return
	}
	fmt.Println(string(data[783]))
	fmt.Println(string(data[784]))
	fmt.Println(string(data[785]))

}
