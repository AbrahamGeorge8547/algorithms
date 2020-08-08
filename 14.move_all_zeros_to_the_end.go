package main

import (
	"fmt"
	"os"
)

// given an array of int move all zeros to end while maintaing relative order of the array.

func main() {
	x := []int{6, 0, 8, 2, 4, 0, 3, 0, 1}
	switch os.Args[1] {
	case "naive":
		{
			// if the current element is non zero we put it at the next available slot. after all elememts processed fill the rest with zeros.
			naive(x)
		}
	case "optimized":
		{
			// pivot is zero and if non zero element is enxounterd swap the elements and increment j
			optimized(x)
		}
	}
}

func naive(x []int) {
	k := 0
	for i := 0; i < len(x); i++ {
		if x[i] != 0 {
			x[k] = x[i]
			k++
		}
	}
	for i := k; i < len(x); i++ {
		x[i] = 0
	}
	fmt.Print(x)
}

func optimized(x []int) {
	j := 0
	for i := 0; i < len(x); i++ {
		if x[i] != 0 {
			swap(x, i, j)
			j++
		}
	}
	fmt.Print(x)
}

func swap(x []int, i, j int) {
	temp := x[i]
	x[i] = x[j]
	x[j] = temp
}
