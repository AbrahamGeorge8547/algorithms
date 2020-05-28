package summing

import "fmt"

func Sort(arr []int) {

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	zeros := len(arr) - sum
	k := 0
	for zeros > 0 {
		arr[k] = 0
		k++
		zeros--
	}

	for k < len(arr) {
		arr[k] = 1
		k++
	}
	fmt.Println(arr)
}
