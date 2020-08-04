package main

// https://medium.com/@kingrayhan/500-data-structures-and-algorithms-practice-problems-and-their-solutions-b45a83d803f0
// http://www.techiedelight.com/find-duplicate-element-limited-range-array/ referance

// array elements are in the range of 1 - (n-1) where n is length of the array
import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	x := []int{1, 2, 3, 3, 4}
	switch os.Args[1] {

	case "hash":
		{
			defer timeTrack(time.Now(), "hash method")
			hash(x)
		}
	case "sum":
		{
			defer timeTrack(time.Now(), "sorting binary method")
			sum(x)
		}
	default:
		fmt.Println("cases")
	}

}

func hash(arr []int) {
	maps := make(map[int]bool)

	for i := 0; i < len(arr); i++ {

		if _, ok := maps[arr[i]]; ok {
			fmt.Println(arr[i])
		} else {
			maps[arr[i]] = true
		}
	}
}

func sum(arr []int) {
	n := len(arr)
	sum := (n * (n - 1)) / 2
	arrsum := 0

	for i := 0; i < n; i++ {
		arrsum += arr[i]

	}
	fmt.Println(arrsum - sum)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
