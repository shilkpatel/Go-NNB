package main

import (
	"gonum.org/v1/gonum/mat"
)

type activation func(x float64) float64

var ReLu activation = func(x float64) float64 {
	if x < 0 {
		return 0.0
	} else {
		return 1.0
	}
}

type network struct {
	weights             mat.Dense
	bias                mat.Dense
	activation_function []int
}
