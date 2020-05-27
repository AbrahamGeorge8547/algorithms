package multimap

import "fmt"

func PrintAllSubarrays(arr []int) {
	var mapArr = make(map[int]int)
	sum := 0
	pairCounter := 0
	mapArr[0] = -1
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if value, ok := mapArr[sum]; ok {
			fmt.Println(value+1, i)
			pairCounter++
		}
		mapArr[sum] = i
	}
	fmt.Println("Pairs :", pairCounter)
}
