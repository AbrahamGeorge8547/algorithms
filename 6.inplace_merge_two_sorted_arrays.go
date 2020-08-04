package main

import "fmt"

// http://www.techiedelight.com/inplace-merge-two-sorted-arrays/

//  Given two sorted arrays x[] and y[] of size m and n merge elements of x[]with elements of array y[] by maintaining sorted order. i.e. fill x[] with first m smallest elements and fill y with the remaining elements.
//  Conversion should be done inplace without using any other data structure.

// Logic:

// We consider each element of x[] and ignore if it is in correct order(smallest amoung the remaining elements) else we swap it with the smallest element in y[] wich is y[0]. after swaping we move it to the position it is supposed to be in to maintain the sorted order.

func main() {

	x := []int{1, 4, 7, 8, 10}
	y := []int{2, 3, 9}

	merge(x, y)

	fmt.Printf("x: %d", x)
	fmt.Printf("y: %d", y)

}

func merge(x, y []int) {
	m := len(x)
	n := len(y)

	for i := 0; i < m; i++ {
		if x[i] > y[0] {
			temp := x[i]
			x[i] = y[0]
			y[0] = temp
			// y[1..n] is already sorted
			first := y[0]
			k := 1
			for k < n && y[k] < first {
				y[k-1] = y[k]
				k = k + 1
			}
			y[k-1] = first

		}
	}
}
