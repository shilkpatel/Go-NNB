package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type activation func(x float64) float64

var ReLu activation = func(x float64) float64 {
	if x < 0 {
		return 0.0
	} else {
		return x
	}
}

var Linear activation = func(x float64) float64 {
	return x
}

type network struct {
	weights             []mat.Dense
	bias                []mat.Dense
	activation_function []activation
	name                string
}

// needs to initalise weights and biases as random
func New(nodes []int, activation_functions []activation, name string) network {
	var weights []mat.Dense
	var bias []mat.Dense
	var i int
	for i = 0; i < len(nodes)-2; i++ {
		weights = append(weights, *mat.NewDense(nodes[i+1], nodes[i], nil))
	}

	var j int
	for j = 1; j < len(nodes)-2; j++ {
		bias = append(bias, *mat.NewDense(nodes[j], 1, nil))
	}

	return network{weights, bias, activation_functions, name}
}

func test_New(weights []mat.Dense, bias []mat.Dense, activation_f []activation, name string) network {
	return network{weights, bias, activation_f, name}
}

func forward_pass(neural network, vector_input mat.Dense) mat.Dense {
	for i, x := range neural.weights {
		var output mat.Dense

		output.Mul(&vector_input, &x)
		output.Add(&output, &neural.bias[i])

		var x, y int = vector_input.Dims()
		for j := 0; j < y; j++ {
			fmt.Println(j)
			output.Set(x-1, j, neural.activation_function[i](output.At(x-1, j)))
		}
		vector_input = output
	}
	return vector_input
}

// finished but untested
func add_layer(neural network, number_of_notes int, activation_function activation) {
	if len(neural.weights) == 0 {
		neural.weights = append(neural.weights, *mat.NewDense(number_of_notes, number_of_notes, nil))
		neural.bias = append(neural.bias, *mat.NewDense(number_of_notes, 1, nil))
	} else {
		var length int = len(neural.bias)
		row, col := neural.weights[length-1].Dims()
		if row == 0 {
			//prevents error from dims function
			row = 1
		}
		//input of this layer equals output of last layer
		neural.weights = append(neural.weights, *mat.NewDense(col, number_of_notes, nil))
		neural.bias = append(neural.bias, *mat.NewDense(number_of_notes, 1, nil))
	}
}
