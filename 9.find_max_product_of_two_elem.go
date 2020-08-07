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
	max1, max2 := x[0], -^int(0)
	min1, min2 := x[0], ^int(0)
	for i := 1; i < len(x); i++ {
		if x[i] > max1 {
			max2 = max1
			max1 = x[i]
		} else if x[i] > max2 {
			max2 = x[i]
		} else if x[i] < min1 {
			min2 = min1
			min1 = x[i]
		} else if x[i] < min2 {
			min2 = x[i]
		}
	}
	if (max1 * max2) > (min1 * min2) {
		fmt.Printf("max elements are %d and %d", max1, max2)
	} else {
		fmt.Printf("max elements are %d and %d", min1, min2)
	}
}
