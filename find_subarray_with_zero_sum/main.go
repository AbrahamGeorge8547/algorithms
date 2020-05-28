package main

// https://medium.com/@kingrayhan/500-data-structures-and-algorithms-practice-problems-and-their-solutions-b45a83d803f0
// http://www.techiedelight.com/find-sub-array-with-0-sum/  referance
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abrahamgeorge8547/algorithms/find_subarray_with_zero_sum/multimap"
	"github.com/abrahamgeorge8547/algorithms/find_subarray_with_zero_sum/naive"
)

func main() {
	x := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	switch os.Args[1] {
	case "naive":
		{
			defer timeTrack(time.Now(), "naive")
			naive.PrintAllSubarrays(x)

		}
	case "map":
		{
			defer timeTrack(time.Now(), "naive")
			multimap.PrintAllSubarrays(x)

		}
	default:
		fmt.Println("cases")
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
