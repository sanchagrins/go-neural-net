package main

import "math"

// sigDeriv calculates the derivative of the Sigmoid Function
func sigDeriv(x[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i:= 0; i<4; i++{
		result[i][0] = x[i][0] * (1-x[i][0])
	}
	return result
}

// matrixVector is a simple matrix-vector multiplicaiton function (4x3)*(3x1)=(4x1)
func matrixVector(inData[4][3]float64, weights[3][1]float64) [4][1]float64{
	var tmp[4][3]float64
	var result[4][1]float64

	for i := 0; i<4; i++ {
            for j := 0; j<3; j++ {
                tmp[i][j] = inData[i][j]*weights[j][0]
             }
        }

	for i := 0; i<4; i++ {
	    var sum float64
            for j := 0; j<3; j++ {
                sum += tmp[i][j]
             }
	     result[i][0] =  sum
        }
	return result
}

// transposeMatrix transposes an input matrix (4x3) --> (3x4)
func transposeMatrix(matrixA[4][3]float64) [3][4]float64{
	var result[3][4]float64

	for i := 0; i<3; i++ {
		for j := 0; j<4; j++ {
			result[i][j] = matrixA[j][i]
		}
	}
	return result
}

// updateWeights updates the weights using matrix-vector multiplication (3x4)*(4x1) = (3x1)
func updateWeights(weights[3][1]float64, inData[3][4]float64, delta[4][1]float64) [3][1]float64{
	var tmp[3][4]float64
	var result[3][1]float64

	for i := 0; i<3; i++ {
            for j := 0; j<4; j++ {
                tmp[i][j] = inData[i][j]*delta[j][0]
             }
        }

	for i := 0; i<3; i++ {
	    var sum float64
            for j := 0; j<4; j++ {
                sum += tmp[i][j]
             }
	     result[i][0] =  sum + weights[i][0]
	     sum = 0
        }

	return result
}

// subMatrix calculates elementwise subtraction of two input vectors
func subMatrix(matrixA[4][1] float64, matrixB[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i := 0; i<4; i++ {
		for j := 0; j<1; j++ {
			result[i][j] = matrixA[i][j] - matrixB[i][j]
		}
	}

	return result
}

// elementMult calculates elementwise multiplication between two vectors
func elementMult(matrixA[4][1] float64, matrixB[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i := 0; i<4; i++ {
		for j := 0; j<1; j++ {
			result[i][j] = matrixA[i][j] * matrixB[i][j]
		}
	}
	return result
}

// subScalar subtracts a scalar from an input vector
func subScalar(scalar float64, inData[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i:= 0; i<4; i++ {
		for j :=0; j<1; j++ {
			result[i][j] = scalar-inData[i][j]
		}
	}
	return result
}

// sigmoid calculates the Sigmoid Function 1/(1-e^-x)
func sigmoid(inData[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i:=0; i<4; i++ {
		for j := 0; j<1; j++ {
			result[i][j] = 1/(1+math.Exp(-inData[i][j]))
		}
	}
	return result
}
