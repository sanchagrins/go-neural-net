package main

func dot(inData[4][3]float64, weights[3][1]float64) [4][3]float64{
	var result[4][3]float64

	for i := 0; i<4; i++ {
            for j := 0; j<3; j++ {
                //fmt.Println(inData[i][j])
                //fmt.Println(weights[j][0])
                result[i][j] = inData[i][j]*weights[j][0]
             }
        }
	return result
}
