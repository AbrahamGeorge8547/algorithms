package naive

import "fmt"

func PrintAllSubarrays(arr []int) {
	pairCounter := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			if arr[j] == 0 {
				continue
			}
			sum += arr[j]
			if sum == 0 {
				fmt.Println(i, j)
				pairCounter++
				break
			}
		}
	}
	fmt.Println("Pairs:", pairCounter)
}
