package main

// http://www.techiedelight.com/find-pair-with-given-sum-array/  referance
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abrahamgeorge8547/algorithms/find_pair_with_given_sum/hashing"
	"github.com/abrahamgeorge8547/algorithms/find_pair_with_given_sum/naive"
	"github.com/abrahamgeorge8547/algorithms/find_pair_with_given_sum/sorting"
	"github.com/abrahamgeorge8547/algorithms/find_pair_with_given_sum/sortingbinary"
)

func main() {
	x := []int{2, 4, 3, 5, 6, -2, 4, 7, 8, 9}
	sum := 7
	switch os.Args[1] {
	case "naive":
		{
			defer timeTrack(time.Now(), "naive")
			naive.FindPair(x, sum)

		}
	case "sort":
		{
			defer timeTrack(time.Now(), "sorting method")
			sorting.FindPair(x, sum)

		}
	case "hash":
		{
			defer timeTrack(time.Now(), "hash method")
			hashing.FindPair(x, sum)
		}
	case "sortingHash":
		{
			defer timeTrack(time.Now(), "sorting binary method")
			sortingbinary.FindPair(x, sum)
		}
	default:
		fmt.Println("cases")
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
