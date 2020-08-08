package main

import (
	"fmt"
	"os"
)

// http://www.techiedelight.com/find-longest-bitonic-subarray-array/
// Q: find the longest sequence in which first elements are sorted in increasing order then in decreasing order

func main() {
	x := []int{3, 5, 8, 4, 5, 9, 10, 8, 5, 3, 4}
	switch os.Args[1] {
	case "auxilary":
		/* keep to auxilary subarrays  I and D I stores the length of the longest increasing subarray ending at x[i]
		D[] stores the longest decreasing array starting from x[i]
		*/
		{
			auxilary(x)
		}
	case "without_auxilary":
		{
			// fist start from i and end at j record the lenght then start at j and reapeat the process while maintainig the length
			withoutAuxilary(x)
		}
	}
}

func auxilary(x []int) {
	I := make([]int, len(x))
	for i := range I {
		I[i] = 1
	}
	D := make([]int, len(x))
	for i := range D {
		D[i] = 1
	}

	for i := 1; i < len(x); i++ {
		if x[i-1] < x[i] {
			I[i] = I[i-1] + 1
		}
	}

	for j := (len(x) - 2); j >= 0; j-- {
		if x[j] > x[j+1] {
			D[j] = D[j+1] + 1
		}
	}

	lbs_len := 0
	beg, end := 0, 0

	for i := range x {
		if lbs_len < I[i]+D[i]-1 {
			lbs_len = I[i] + D[i] - 1
			beg = i - I[i] + 1
			end = i + D[i] - 1
		}
	}

	fmt.Printf("Beg: %d End: %d", beg, end)

}

func withoutAuxilary(x []int) {
	endIndex := 0
	maxLen := 0
	i := 0
	for i+1 < len(x) {
		length := 1
		for i+1 < len(x) && x[i] < x[i+1] {
			i++
			length++
		}
		for i+1 < len(x) && x[i] > x[i+1] {
			i++
			length++
		}
		if length > maxLen {

			maxLen = length
			endIndex = i
		}
	}
	fmt.Printf("Beg: %d End: %d", endIndex-maxLen+1, endIndex)
}
