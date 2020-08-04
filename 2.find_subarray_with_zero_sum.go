package main

// https://medium.com/@kingrayhan/500-data-structures-and-algorithms-practice-problems-and-their-solutions-b45a83d803f0
// http://www.techiedelight.com/find-sub-array-with-0-sum/  referance
import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	x := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	switch os.Args[1] {
	case "naive":
		{
			// Plain brute force
			defer timeTrack(time.Now(), "naive")
			naive(x)

		}
	case "map":
		{
			defer timeTrack(time.Now(), "naive")
			multimap(x)

		}
	default:
		fmt.Println("cases")
	}

}

func naive(arr []int) {
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

func multimap(arr []int) {

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

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
