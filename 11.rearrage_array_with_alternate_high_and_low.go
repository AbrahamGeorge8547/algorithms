package main

import "fmt"

// http://www.techiedelight.com/rearrange-the-array-with-alternate-high-and-low-elements/
// Given an array of integers, rearrange the array such that every second element of the array is greater than its left and right elements. Assume no duplicate elements are present in the array... A simple solution would be to sort the array in ascending order first. Then we take another auxiliary array and fill it with..

// An efficient solution doesn’t involve sorting the array or use of auxiliary space. We start from the second element of the array and increment index by 2 for each iteration of loop. If previous element is greater than the current element, we swap the elements. Similarly if next element is greater than the current element, we swap both elements.

func main() {
	x := []int{9, 6, 8, 3, 7, 10, 14, 16}

	for i := 1; i < len(x); i += 2 {
		if x[i-1] > x[i] {
			swap(x, i-1, i)
		}

		if i+1 < len(x) && x[i+1] > x[i] {
			swap(x, i+1, i)
		}
	}

	fmt.Printf("%d", x)
}

func swap(x []int, i, j int) {
	temp := x[i]
	x[i] = x[j]
	x[j] = temp
}
