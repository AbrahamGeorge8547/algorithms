package main

import "fmt"

// http://www.techiedelight.com/merge-two-arrays-satisfying-given-constraints/

// give two sorted arrays x[] and y[] of size m and n each, m> n.( Vacant cells in x represented by 0)
// x[] has exactly n vacant cells merge elemets of y into x in their correct position
// Answer: Move all elements of x to the begining.
func main() {

	X := []int{0, 2, 0, 3, 0, 5, 6, 0, 0}
	Y := []int{1, 4, 9, 10, 15}

	rearrage(X)

	merge(X, Y)

	fmt.Print(X)
}

func rearrage(x []int) {

	// moving all elements of x to beggining
	k := 0
	for i := 0; i < len(x); i++ {

		if x[i] > 0 {
			x[k] = x[i]
			x[i] = 0
			k = k + 1
		}
	}
}

func merge(x, y []int) {
	k := len(x) - 1
	n := len(y) - 1
	m := k - n - 1
	for m >= 0 && n >= 0 {
		if x[m] > y[n] {
			x[k] = x[m]
			m = m - 1
		} else {
			x[k] = y[n]
			n = n - 1
		}
		k = k - 1
	}

	for n >= 0 {
		x[k] = y[n]
		n = n - 1
	}
}
