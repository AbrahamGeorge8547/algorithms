package sortingbinary

import (
	"errors"
	"fmt"
	"sort"
)

func FindPair(arr []int, sum int) {
	sort.Ints(arr)
	pairs := 0
	for i := 0; i < len(arr); i++ {
		if index, err := binarySearch(sum-arr[i], arr); err == nil {
			fmt.Println(index, arr[i])
			pairs++
		}
	}
	fmt.Println("pairs: ", pairs)
}

func binarySearch(elem int, arr []int) (int, error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2
		if arr[median] < elem {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != elem {
		return 0, errors.New("not found")
	}
	return arr[low], nil
}
