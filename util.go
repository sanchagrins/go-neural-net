package main

import "math"
//import "fmt"

func dot(x[4][3]float64, y[3][1]float64) [4][1]float64{
	var result[4][1]float64

	for i := 0; i<4; i++ {
		for j := 0; j<3; j++ {
			result[i][j] = x[i][j] * y[i][j]
		}
	}
	return result
}

func sigDiv(x[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i:= 0; i<4; i++{
		result[i][0] = x[i][0] * (1-x[i][0])
	}
	return result
}

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

func transposeMatrix(matrixA[4][3]float64) [3][4]float64{
	var result[3][4]float64

	for i := 0; i<3; i++ {
		for j := 0; j<4; j++ {
			result[i][j] = matrixA[j][i]
		}
	}
	return result
}

func updateWeights(weights[3][1]float64, inData[3][4]float64, delta[4][1]float64) [3][1]float64{
	var tmp[3][4]float64
	var result[3][1]float64

	for i := 0; i<3; i++ {
            for j := 0; j<4; j++ {
                tmp[i][j] = inData[i][j]*delta[j][0]
             }
        }
	//fmt.Println("l0.T*l1_delta: ", tmp)

	for i := 0; i<3; i++ {
	    var sum float64
            for j := 0; j<4; j++ {
                sum += tmp[i][j]
             }
	     result[i][0] =  sum + weights[i][0]
	     sum = 0
        }

	//fmt.Println("l0.T dot l1_delta =", result);
	return result
}

func subMatrix(matrixA[4][1] float64, matrixB[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i := 0; i<4; i++ {
		for j := 0; j<1; j++ {
			result[i][j] = matrixA[i][j] - matrixB[i][j]
		}
	}

	return result
}

func elementMult(matrixA[4][1] float64, matrixB[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i := 0; i<4; i++ {
		for j := 0; j<1; j++ {
			result[i][j] = matrixA[i][j] * matrixB[i][j]
		}
	}
	return result
}

func subScalar(scalar float64, inData[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i:= 0; i<4; i++ {
		for j :=0; j<1; j++ {
			result[i][j] = scalar-inData[i][j]
		}
	}
	return result
}

func sigmoid(inData[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i:=0; i<4; i++ {
		for j := 0; j<1; j++ {
			result[i][j] = 1/(1+math.Exp(-inData[i][j]))
		}
	}
	return result
}
