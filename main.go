package main

import "fmt"
import "math/rand"

// Sigmoid Function
func nonlin(x[4][1]float64, deriv bool) [4][1]float64 {
	if deriv == true {
		return dot(x, subScalar(1.0, x))
	} else {
		return sigmoid(x)
	}
}

func main() {

	// Input Dataset
	var inData =[4][3]float64{
		{0,0,1},
		{0,1,1},
		{1,0,1},
		{1,1,1},
	}

	var outData =[4][1]float64{
		{0},{0},{1},{1},
	}

	// Seed random numbers to make calculations
	// deterministic
	rand.Seed(1)

	// Initialize wieghts randomly with mean 0
	var syn0 = [3][1]float64{
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
	 }


	var result[4][1]float64
	var deriv bool = true

	for i:=0; i<=10; i++ {

		// Forward Propagation
		l0 := inData
		l1 := nonlin(matrixVector(l0, syn0), false)
		fmt.Println("l1 = nonlin(l0 * syn0): ", l1)

		// Calculate Error
		l1_error := subMatrix(outData, l1)
		fmt.Println("l1_error: ", l1_error)

		// Multiply how error by the derivative of the sigmoid of l1
		l1_delta := elementMult(l1_error, nonlin(l1, true))
		fmt.Println("l1_delta: ", l1_delta)

		// Update weights
		//syn0 += matrixVector(l0, l1_delta)
	}

	fmt.Println("outData", outData)

	result = matrixVector(inData, syn0)
	fmt.Println("Resutl of matrixVector", result)

	result = nonlin(result, deriv)
	fmt.Println("Result of nonlin", result)

	deriv = false
	result = nonlin(result, deriv)
	fmt.Println("Result of nonlin", result)
}
