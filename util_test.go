package main

import "testing"

func Test_subScalar(t *testing.T) {
	v := [][]float64{{1},{1},{1},{1}}

	val := subScalar(1,v)

	for i:=0;i<len(val);i++ {
		if val[i][0] != 0 {
			t.Fatalf("Error: Expected '0', found '%v'", val[i][0])
		}
	}
}

func Test_sigmoid(t *testing.T) {
	v := [][]float64{{0},{0},{0},{0}}

	val := sigmoid(v)

	for i:=0;i<len(val);i++ {
		if val[i][0] != 0.5 {
			t.Fatalf("Error: Expected '0.5', found '%v'", val[i][0])
		}
	}
}

func Test_elementMult(t *testing.T) {
	v := [][]float64{{2},{2},{2},{2}}

	val := elementMult(v,v)

	for i:=0;i<len(val);i++ {
		if val[i][0] != 4 {
			t.Fatalf("Error: Expected '4', found '%v'", val[i][0])
		}

	}
}
