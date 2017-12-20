package main

import "testing"

func Test_subScalar(t *testing.T) {
	var v = [4][1]float64{
		{1},{1},{1},{1},
	}

	val := subScalar(1,v)

	for i:=0;i<4;i++ {
		if val[i][0] != 0 {
			t.Fatalf("Error: Expected '0', found '%v'", val[i][0])
		}
	}
}

