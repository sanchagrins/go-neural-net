package main

import "fmt"
import "math/rand"

func main() {
	var inData =[4][3]float64{
		{0,0,1},
		{0,1,1},
		{1,0,1},
		{1,1,1},
	}
	var outData =[1][4]int{
		{0,0,1,1},
	}

	rand.Seed(1)

	var weights = [3][1]float64{
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
	 }
	var result[4][3]float64

	fmt.Println(inData)
	fmt.Println(outData)

	fmt.Println(weights)
	for i:=0; i<=10000; i++ {
		//l0 := inData
	}

	result = dot(inData, weights)
	fmt.Println(result)
}
