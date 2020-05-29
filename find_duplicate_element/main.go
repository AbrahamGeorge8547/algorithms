package main

// https://medium.com/@kingrayhan/500-data-structures-and-algorithms-practice-problems-and-their-solutions-b45a83d803f0
// http://www.techiedelight.com/find-duplicate-element-limited-range-array/ referance

// array elements are in the range of 1 - (n-1) where n is length of the array
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abrahamgeorge8547/algorithms/find_duplicate_element/hash"
	"github.com/abrahamgeorge8547/algorithms/find_duplicate_element/sum"
)

func main() {
	x := []int{1,2,3,3,4}
	switch os.Args[1] {

	case "hash":
		{
			defer timeTrack(time.Now(), "hash method")
			hash.Findduplicate(x)
		}
	case "sum":
		{
			defer timeTrack(time.Now(), "sorting binary method")
			sum.Findduplicate(x)
		}
	default:
		fmt.Println("cases")
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
