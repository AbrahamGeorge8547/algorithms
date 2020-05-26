package sorting

import (
	"fmt"
	"sort"
)

func FindPair(arr []int, sum int) {
	sort.Ints(arr)
	low := 0
	high := len(arr) - 1
	k := 0
	fmt.Println(arr)
	for low < high {
		if arr[low]+arr[high] == sum {
			fmt.Println(arr[low], arr[high])
			k++
		}
		if arr[low]+arr[high] < sum {
			low++
		} else {
			high--
		}
	}
	fmt.Println("pairs:", k)
	// sortedArr := sort.Ints(arr[:])
	// fmt.Print(sortedArr)
}
