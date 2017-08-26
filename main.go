package main

import "fmt"
import "math/rand"
import "gonum.org/v1/gonum/mat"

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

	rand.Seed(1)
	//var syn0 = rand.Float64() * (0)-1
	var weights = [3][1]float64{
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
		 {(rand.Float64()*2)-1},
	 }

	fmt.Println(inData)
	fmt.Println(outData)

	fmt.Println(weights)
	for i:=0; i<=10000; i++ {
		l0 := inData
	}
}
