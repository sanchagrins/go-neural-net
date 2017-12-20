package main

import "fmt"
import "math/rand"

// nonlin calculates the Sigmoid Function and derivative for a given matrix
func nonlin(x[][]float64, deriv bool) [][]float64 {
	if deriv == true {
		return sigDeriv(x)
	} else {
		return sigmoid(x)
	}
}

func main() {
	// First Layer of the nextwork specified by the input data, a 4x3 matrix where each row is a training example
	var l0 =[][]float64{
		{0,0,1},
		{0,1,1},
		{1,0,1},
		{1,1,1},
	}

	// Output 4x1 matrix where each row is a training example
	var outData =[][]float64{
		{0},{0},{1},{1},
	}

	// Second layer of the network (Hidden Layer)
	var l1[][]float64

	// Seed random numbers to make calculations deterministic
	rand.Seed(1)

	// First layer of weights. 3x1 vector with values generated randomly between -1 and 1 with mean 0 
	var syn0 = [][]float64{
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
	 }

	// Transpose the input matrix. Used for updating the wieghts
	l0T := transposeMatrix(l0)

	for i:=0; i<100000; i++ {

		// Forward Propagation
		// Sigmoid Function calculated with the product the input matrix and weights vector
		l1 = nonlin(matrixVector(l0, syn0), false)

		// Calculate Error: Elementwise vector subtraction 
		l1_error := subMatrix(outData, l1)

		// Calculate Error Weighted Derivative: Elementwise multiplication with the l1_error and the derivative of
		// the Sigmoid function for the hidden layer
		l1_delta := elementMult(l1_error, nonlin(l1, true))

		// Update the weights
		syn0 = updateWeights(syn0, l0T, l1_delta)
	}

	fmt.Println("Results after training: ", l1)
}
