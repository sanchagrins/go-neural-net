package main

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

func subScalar(scalar float64, inData[4][1]float64) [4][1]float64{
	var result[4][1]float64

	for i:= 0; i<4; i++ {
		for j :=0; j<1; j++ {
			result[i][j] = scalar-inData[i][j]
		}
	}
	return result
}
