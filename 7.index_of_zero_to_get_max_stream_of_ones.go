package main

import "fmt"

// http://www.techiedelight.com/find-index-0-replaced-get-maximum-length-sequence-of-continuous-ones/

//  given a binary array find the index of zero to be replaced by 1 to get the maximum length sequenxce of ones

//  Ans: update each non zero elements of array with count of  their adjecent consecutive ones.  then dfind the index of zero with max sum on left and right.

func main() {

	x := []int{0, 0, 1, 0, 1, 1, 1, 0, 1, 1}
	index := findIndex(x)
	fmt.Print(index)
}

func findIndex(x []int) int {

	for i := 1; i < len(x); i++ {
		if x[i] == 1 {
			x[i] += x[i-1]
		}
	}
	count := 0

	for i := len(x) - 1; i >= 0; i-- {
		if x[i] > count {
			count = x[i]
		}
		if x[i] > 0 {
			x[i] = count
		} else {
			count = 0
		}
	}

	maxCount := 0
	maxIndex := -1

	for i := 0; i < len(x); i++ {
		if x[i] == 0 {

			if i == 0 && maxCount < x[i+1]+1 {
				maxCount = x[i+1] + 1
				maxIndex = i
			} else if i == len(x)-1 && maxCount < x[i-1]+1 {
				maxCount = x[i-1] + 1
				maxIndex = i
			} else if maxCount < x[i-1]+x[i+1]+1 {
				maxCount = x[i-1] + x[i+1] + 1
				maxIndex = i
			}

		}
	}
	return maxIndex
}
