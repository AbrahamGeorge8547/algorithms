package main

import (
	"fmt"
	"os"
)

// www.techiedelight.com/find-equilibrium-index-array/

// Q: given an array of integers find equilibrium index in it

func main() {
	x := []int{0, -3, 5, -4, -2, 3, 1, 0}
	switch os.Args[1] {
	// naive solution is to calculat sum to the right and left if both are same we have equilibrium
	case "linear":
		/* fist make an array left with sum from x[0..i] in left[i] then after it is filled we traverse the array f
		rom right to left with sum from right and compare it with left[i] and if found to be same then we have equilibrium
		at index 0 */
		{
			var left = make([]int, len(x))
			left[0] = 0
			for i := 1; i < len(x); i++ {
				left[i] = left[i-1] + x[i-1]
			}
			fmt.Print(len(left))
			right := 0

			for i := (len(x) - 1); i >= 0; i-- {
				if left[i] == right {
					fmt.Printf("Equilibrium index found at %d", i)
				}
				right += x[i]
			}
		}
	case "optimized":
		//  sum of left array(0...i-1) = total - (x[i]+ sum of right array(i+1...n))
		{
			right := 0
			total := sum(x)

			for i := (len(x) - 1); i >= 0; i-- {
				if right == total-(x[i]+right) {
					fmt.Printf("Equlibrium found at index %d", i)
				}
				right += x[i]
			}

		}
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
