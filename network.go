package main

import (
	"gonum.org/v1/gonum/mat"
)

type activation func(x float64) float64

var ReLu = func(x float64) float64 {
	if x < 0 {
		return 0.0
	} else {
		return 1.0
	}
}

type network struct {
	weights             []mat.Dense
	bias                []mat.Dense
	activation_function []activation
}

// needs to initalise weights and biases as random
func New(nodes []int, activation_functions []activation) network {
	var weights []mat.Dense
	var bias []mat.Dense
	var i int
	for i = 0; i < len(nodes)-2; i++ {
		weights = append(weights, *mat.NewDense(nodes[i], nodes[i+1], nil))
	}

	var j int
	for j = 1; j < len(nodes)-2; j++ {
		bias = append(bias, *mat.NewDense(nodes[j], 1, nil))
	}

	return network{weights, bias, activation_functions}
}

func Map()

func forward_pass(neural network, vector_input mat.Dense) mat.Dense {
	for i, x := range neural.weights {
		vector_input.Mul(&vector_input, &x)
		vector_input.Add(&vector_input, &neural.bias[i])
		var x, y int = vector_input.Dims()

		for j := 0; j < x; j++ {
			vector_input.Set(j, y, neural.activation_function[i](vector_input.At(j, y)))
		}
	}
	return vector_input
}
