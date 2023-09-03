package main

import "fmt"

func main() {

	var x int
	x = 10
	fmt.Println(x)

	var y [10]int
	var t = [2]int{1, 3}
	var slice = []int{10, 30, 40}
	for i := 0; i < 10; i++ {
		y[i] = 10
		fmt.Println(y[i])
	}
	for k := 0; k < len(slice); k++ {
		fmt.Printf("Value of Slices : %d\n", slice[k])
	}
	for j := 0; j < len(t); j++ {
		fmt.Printf("Value of T : %d\n", t[j])
	}

	z := [2][2]int{
		{1, 2},
		{1, 2},
	}
	for c := 0; c < 2; c++ {
		fmt.Println(z[c][0])
	}
}
