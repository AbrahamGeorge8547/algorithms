package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	x := []int{0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1}
	switch os.Args[1] {
	case "sum":
		{
			// find the sum of the array ie no of ones then insert len - sum zeros and rest ones.
			defer timeTrack(time.Now(), "sum")
			summing(x)

		}
	case "next":
		{
			// if the current element is zero the following next available element is put zero after traversing the array rest of the elements are put as one.
			defer timeTrack(time.Now(), "next method")
			next(x)
		}
	default:
		fmt.Println("cases")
	}

}

func summing(arr []int) {
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

func next(arr []int) {
	k := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[k] = 0
			k++
		}
	}
	for j := k; j < len(arr); j++ {
		arr[j] = 1
	}
	fmt.Println(arr)
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
