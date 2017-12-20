package main

import "math"

// sigDeriv calculates the derivative of the Sigmoid Function
func sigDeriv(x [][]float64) [][]float64{
	result := make([][]float64, len(x))
	for i := range result {
		result[i] = make([]float64, len(x[0]))
	}

	for i:= 0; i<len(x); i++{
		result[i][0] = x[i][0] * (1-x[i][0])
	}
	return result
}

// matrixVector is a simple matrix-vector multiplicaiton function (4x3)*(3x1)=(4x1)
func matrixVector(inData[][]float64, weights[][]float64) [][]float64{
	tmp := make([][]float64, len(inData))
	for i := range tmp {
		tmp[i] = make([]float64, len(inData))
	}

	result := make([][]float64, len(inData))
	for i := range result {
		result[i] = make([]float64, len(weights))
	}

	for i := 0; i<len(inData); i++ {
            for j := 0; j<len(inData[0]); j++ {
                tmp[i][j] = inData[i][j]*weights[j][0]
             }
        }

	for i := 0; i<len(inData); i++ {
	    var sum float64
            for j := 0; j<len(inData[0]); j++ {
                sum += tmp[i][j]
             }
	     result[i][0] =  sum
        }
	return result
}

// transposeMatrix transposes an input matrix (4x3) --> (3x4)
func transposeMatrix(matrixA[][]float64) [][]float64{

	result := make([][]float64, len(matrixA[0]))
	for i := range result {
		result[i] = make([]float64, len(matrixA))
	}

	for i := 0; i<len(matrixA[0]); i++ {
		for j := 0; j<len(matrixA); j++ {
			result[i][j] = matrixA[j][i]
		}
	}
	return result
}

// updateWeights updates the weights using matrix-vector multiplication (3x4)*(4x1) = (3x1)
func updateWeights(weights[][]float64, inData[][]float64, delta[][]float64) [][]float64{
	tmp :=  make([][]float64, len(inData))
	for i := range tmp {
		tmp[i] = make([]float64, len(inData[0]))
	}

	result := make([][]float64, len(weights))
	for i := range result {
		result[i] = make([]float64, len(weights[0]))
	}

	for i := 0; i<len(inData); i++ {
            for j := 0; j<len(inData[0]); j++ {
                tmp[i][j] = inData[i][j]*delta[j][0]
             }
        }

	for i := 0; i<len(inData); i++ {
	    var sum float64
            for j := 0; j<len(inData[0]); j++ {
                sum += tmp[i][j]
             }
	     result[i][0] =  sum + weights[i][0]
	     sum = 0
        }

	return result
}

// subMatrix calculates elementwise subtraction of two input vectors
func subMatrix(matrixA[][]float64, matrixB[][]float64) [][]float64{
	result := make([][]float64, len(matrixA))
	for i := range result {
		result[i] = make([]float64, len(matrixA[0]))
	}

	for i := 0; i<len(matrixA); i++ {
		for j := 0; j<len(matrixA[0]); j++ {
			result[i][j] = matrixA[i][j] - matrixB[i][j]
		}
	}

	return result
}

// elementMult calculates elementwise multiplication between two vectors
func elementMult(matrixA[][]float64, matrixB[][]float64) [][]float64{
	result := make([][]float64, len(matrixA))
	for i:= range result {
		result[i] = make([]float64, len(matrixA[0]))
	}

	for i := 0; i<len(matrixA); i++ {
		for j := 0; j<len(matrixA[0]); j++ {
			result[i][j] = matrixA[i][j] * matrixB[i][j]
		}
	}
	return result
}

// subScalar subtracts a scalar from an input vector
func subScalar(scalar float64, inData[][]float64) [][]float64{
	result := make([][]float64, len(inData))
	for i := range(inData[0]) {
		result[i] = make([]float64, len(inData[0]))
	}

	for i:= 0; i<len(inData); i++ {
		for j :=0; j<len(inData[0]); j++ {
			result[i][j] = scalar-inData[i][j]
		}
	}
	return result
}

// sigmoid calculates the Sigmoid Function 1/(1-e^-x)
func sigmoid(inData[][]float64) [][]float64{
	result := make([][]float64, len(inData))
	for i := range result {
		result[i] = make([]float64, len(inData[0]))
	}

	for i:=0; i<len(inData); i++ {
		for j := 0; j<len(inData[0]); j++ {
			result[i][j] = 1/(1+math.Exp(-inData[i][j]))
		}
	}
	return result
}
