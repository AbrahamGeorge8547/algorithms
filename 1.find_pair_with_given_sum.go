package main

// https://medium.com/@kingrayhan/500-data-structures-and-algorithms-practice-problems-and-their-solutions-b45a83d803f0
// http://www.techiedelight.com/find-pair-with-given-sum-array/  referance
import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	x := []int{2, 4, 3, 5, 6, -2, 4, 7, 8, 9}
	sum := 7
	switch os.Args[1] {
	case "naive":
		{
			defer timeTrack(time.Now(), "naive")
			//  consider every pair and return if sum is found.
			naive(x, sum)

		}
	case "sort":
		{
			// The idea is to sort the given array in ascending order and maintain search space by maintaining two indices (low and high) that initially points to two end-points of the array. Then we loop till low is less than high index and reduce search space
			//  at each iteration of the loop. We compare sum of elements present at index low and high with desired sum and increment low if sum is less than the desired sum else we decrement high is sum is more than the desired sum. Finally, we return if pair found in the array.
			defer timeTrack(time.Now(), "sorting method")
			sorting(x, sum)

		}
	case "hash":
		{
			// the idea to insert all elements to a map. if sum - arr[i] exits they are a pair. then add it to the map.
			defer timeTrack(time.Now(), "hash method")
			hashing(x, sum)
		}
	case "sortingHash":
		{
			defer timeTrack(time.Now(), "sorting binary method")
			sortingbinary(x, sum)
		}
	default:
		fmt.Println("cases")
	}

}

func naive(arr []int, sum int) {
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

func sorting(arr []int, sum int) {
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
}

func sortingbinary(arr []int, sum int) {
	sort.Ints(arr)
	pairs := 0
	for i := 0; i < len(arr); i++ {
		if index, err := binarySearch(sum-arr[i], arr); err == nil {
			fmt.Println(index, arr[i])
			pairs++
		}
	}
	fmt.Println("pairs: ", pairs)
}

func hashing(arr []int, sum int) {
	var mapArr = make(map[int]int)
	pairs := 0
	for i := 0; i < len(arr); i++ {
		if value, ok := mapArr[sum-arr[i]]; ok {
			println(value, i)
			pairs++
		}

		mapArr[arr[i]] = i
	}
	println("pairs:", pairs)
}

func binarySearch(elem int, arr []int) (int, error) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2
		if arr[median] < elem {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != elem {
		return 0, errors.New("not found")
	}
	return arr[low], nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
