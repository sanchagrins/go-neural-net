package main

import "fmt"
import "math/rand"

func main() {
	var inData =[4][3]int{
		{0,0,1},
		{0,1,1},
		{1,0,1},
		{1,1,1},
	}
	var outData =[1][4]int{
		{0,0,1,1},
	}

	//rand.Seed(1)
	//var syn0 = rand.Float64() * (0)-1
	var weights = [4][1]float64{
		 {(rand.Float64()*1)-1},
		 {(rand.Float64()*1)-1},
		 {(rand.Float64()*1)-1},
		 {(rand.Float64()*1)-1},
	 }

	fmt.Println(inData)
	fmt.Println(outData)

	fmt.Println(weights)
}
