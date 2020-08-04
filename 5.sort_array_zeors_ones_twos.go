package main

import "fmt"

//  values less than, greater than and equal to  pivot
//  http://www.techiedelight.com/sort-array-containing-0s-1s-2s-dutch-national-flag-problem/
func main() {

	A := []int{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0}
	threeWayPartition(A)
	fmt.Print(A)
}

func swap(x []int, i, j int) {
	temp := x[i]
	x[i] = x[j]
	x[j] = temp
}

func threeWayPartition(x []int) {

	end := len(x) - 1

	start, mid := 0, 0
	pivot := 1

	for mid <= end {
		if x[mid] < pivot {
			swap(x, start, mid)
			start = start + 1
			mid = mid + 1
		} else if x[mid] > pivot {
			swap(x, mid, end)
			end = end - 1
		} else {
			mid = mid + 1
		}
	}

}
