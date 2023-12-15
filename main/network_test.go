package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestRelu(t *testing.T) {
	assert.Equal(t, 3.0, ReLu(3))
	assert.Equal(t, 0.0, ReLu(0))
	assert.Equal(t, 0.0, ReLu(-5))
}

func TestNetwork_1(t *testing.T) {
	weight := mat.NewDense(2, 2, []float64{1, 1, 1, 1})
	bias := mat.NewDense(1, 2, []float64{3, 4})
	test_network := test_New([]mat.Dense{*weight}, []mat.Dense{*bias}, []activation{Linear})
	input := mat.NewDense(1, 2, []float64{0, 1})
	result := forward_pass(test_network, *input)
	assert.Equal(t, 4.0, result.At(0, 0))
	assert.Equal(t, 5.0, result.At(0, 1))
}

func TestNetwork_2(t *testing.T) {

}
