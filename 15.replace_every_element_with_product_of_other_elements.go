package main

import (
	"fmt"
	"os"
)

// http://www.techiedelight.com/replace-element-array-product-every-element-without-using-division-operator/

/* Q: given an array of integers, replace every ekement of the array with the product
of every other element of the array without using the division operator.
*/

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Print(os.Args[1])

	switch os.Args[1] {
	case "auxilary":
		{

			// use two auxilary array left and right which stores product of 0..i-1 and i+1..n and multiply left[i] * right[i]
			auxilary(x)
		}

	case "recursive":
		{
			recursive(x, 1, 0)
			fmt.Printf("Output: %d", x)
		}

	}
}

func auxilary(x []int) {
	left := make([]int, len(x))
	right := make([]int, len(x))
	left[0] = 1
	right[len(x)-1] = 1
	for i := 1; i < len(x); i++ {
		left[i] = left[i-1] * x[i-1]
	}

	for j := (len(x) - 2); j >= 0; j-- {
		right[j] = right[j+1] * x[j+1]
	}
	for i := 0; i < len(x); i++ {
		x[i] = left[i] * right[i]
	}

	fmt.Print(x)
}

func recursive(x []int, left, i int) int {
	if len(x) == i {
		return 1
	}
	curr := x[i]
	right := recursive(x, left*x[i], i+1)
	x[i] = left * right
	return curr * right

}
