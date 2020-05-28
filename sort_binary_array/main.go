package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abrahamgeorge8547/algorithms/sort_binary_array/next"
	"github.com/abrahamgeorge8547/algorithms/sort_binary_array/summing"
)

func main() {
	x := []int{0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1}
	switch os.Args[1] {
	case "sum":
		{
			defer timeTrack(time.Now(), "sum")
			summing.Sort(x)

		}
	case "next":
		{
			defer timeTrack(time.Now(), "next method")
			next.Sort(x)
		}
	default:
		fmt.Println("cases")
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
