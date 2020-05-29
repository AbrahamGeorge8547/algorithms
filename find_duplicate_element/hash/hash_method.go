package hash

import "fmt"

func Findduplicate(arr []int) {
	maps := make(map[int]bool)

	for i := 0; i < len(arr); i++ {

		if _, ok := maps[arr[i]]; ok {
			fmt.Println(arr[i])
		} else {
			maps[arr[i]] = true
		}
	}
}
