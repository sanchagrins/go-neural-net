package main

import "fmt"
import "math/rand"

// Sigmoid Function
func nonlin(x[4][1]float64, deriv bool) [4][1]float64 {
	if deriv == true {
		return sigDiv(x)
	} else {
		return sigmoid(x)
	}
}

func main() {
	var l1[4][1]float64
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
	  fmt.Println("Original Weight: ", syn0)

	for i:=0; i<100000; i++ {

		// Forward Propagation
		l0 := inData
		tmp := matrixVector(l0, syn0)
		//fmt.Println("l0 dot syn0", tmp)

		l1 = nonlin(tmp, false)
		//fmt.Println("l1 = nonlin(l0 * syn0): ", l1)

		// Calculate Error
		l1_error := subMatrix(outData, l1)
		fmt.Println("l1_error = outdata - l1 ", l1_error)

		// Multiply how error by the derivative of the sigmoid of l1
		//fmt.Println("l1_diriv: ", nonlin(l1, true))
		l1_delta := elementMult(l1_error, nonlin(l1, true))
		//fmt.Println("l1_delta = l1_error * nonlin(li) ", l1_delta)

		//fmt.Println("l0: ", l0)
		//fmt.Println("transpose(l0): ", transposeMatrix(l0))
		// Update weights
		//fmt.Println("l0 = ", l0)
		l0T := transposeMatrix(l0)
		//fmt.Println("l0T = ", l0T)

		syn0 = updateWeights(syn0, l0T, l1_delta)
		fmt.Println("Updated weight: ", syn0)
	}

	//fmt.Println("outData", outData)

	//result = matrixVector(inData, syn0)
	//fmt.Println("Resutl of matrixVector", result)

	//result = nonlin(result, deriv)
	//fmt.Println("Result of nonlin", result)

	//deriv = false
	//result = nonlin(result, deriv)
	//fmt.Println("Result of nonlin", result)
	fmt.Println("Results after training: ", l1)
}
