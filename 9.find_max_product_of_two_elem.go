package main

import (
	"fmt"
	"os"
)

// http://www.techiedelight.com/find-maximum-product-two-integers-array/

// another option is to sort the array and then the max is formed by the product of first two elements or last two elements.

func main() {
	x := []int{-10, -3, 5, -7, -2}
	switch os.Args[1] {
	case "naive":
		{
			//  brute force
			naive(x)
		}

	case "max_min":
		// find max, second max, min, second min
		{
			max_min(x)
		}
	}
}

func naive(x []int) {
	maxint := -^int(0)
	maxI, maxJ := -1, -1
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if maxint < x[i]*x[j] {
				maxint = x[i] * x[j]
				maxI, maxJ = i, j
			}
		}
	}
	fmt.Printf("%d%d", maxI, maxJ)
}

func max_min(x []int) {
	max_1, max_2 := x[0], -^int(0)
	min_1, min_2 := x[0], ^int(0)
	for i := 1; i < len(x); i++ {
		if x[i] > max_1 {
			max_2 = max_1
			max_1 = x[i]
		} else if x[i] > max_2 {
			max_2 = x[i]
		} else if x[i] < min_1 {
			min_2 = min_1
			min_1 = x[i]
		} else if x[i] < min_2 {
			min_2 = x[i]
		}
	}
	if (max_1 * max_2) > (min_1 * min_2) {
		fmt.Printf("max elements are %d and %d", max_1, max_2)
	} else {
		fmt.Printf("max elements are %d and %d", min_1, min_2)
	}
}
