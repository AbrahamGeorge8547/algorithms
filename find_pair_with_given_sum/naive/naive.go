package naive

import (
	"fmt"
)

func FindPair(arr []int, sum int) {
	var i, j int
	k := 0
	for i = 0; i < len(arr)-1; i++ {
		for j = i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Println(i, j)
				k++
			}
		}
	}
	fmt.Println("pairs:", k)
}
