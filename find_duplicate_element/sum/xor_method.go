package sum

import "fmt"

func Findduplicate(arr []int) {
	n := len(arr)
	sum := (n * (n - 1)) / 2
	arrsum := 0

	for i := 0; i < n; i++ {
		arrsum += arr[i]

	}
	fmt.Println(arrsum - sum)
}
