package main

import (
	"fmt"
	"math/rand"
	"time"
)

// http://www.techiedelight.com/shuffle-given-array-elements-fisher-yates-shuffle/

// given an array of integers, inplace shuffle i. The algorithm should produce an unbiased permutaiton. every permutaiton must be equally likely
func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}

	shuffle(x)

}

func shuffle(x []int) {
	rand.Seed(time.Now().UTC().UnixNano())
	length := len(x)
	for i := 0; i < length-1; i++ {
		j := rand.Intn(length-i) + i
		swap(x, i, j)
	}
	fmt.Printf("%d", x)
}

func swap(x []int, i, j int) {
	temp := x[i]
	x[i] = x[j]
	x[j] = temp
}
