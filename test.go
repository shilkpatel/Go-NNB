package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	fmt.Println("Hello world")
	a := 4
	print(a)
	zero := mat.NewDense(2, 2, nil)
	print(zero)

}
