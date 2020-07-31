package main

import (
	"max_product/max_min"
	"max_product/naive"
	"os"
)

// http://www.techiedelight.com/find-maximum-product-two-integers-array/

// another option is to sort the array and then the max is formed by the product of first two elements or last two elements.

func main() {
	x := []int{-10, -3, 5, -7, -2}
	switch os.Args[1] {
	case "naive":
		{
			naive.Find(x)
		}

	case "max_min":
		// find max, second max, min, second min
		{
			max_min.Find(x)
		}
	}
}
