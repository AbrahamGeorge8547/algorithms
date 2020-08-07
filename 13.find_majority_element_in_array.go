package main

import (
	"fmt"
	"os"
)

// find the majority element in the array
//  majority element if element present more than n/2 times where n is the size

// http://www.techiedelight.com/find-majority-element-in-an-array-boyer-moore-majority-vote-algorithm/

func main() {
	x := []int{2, 2, 3, 1, 4, 1, 7, 4, 2, 2, 2, 2}

	switch os.Args[1] {
	case "naive":
		{
			//  find the frequency of elements in the first half of the array.
			index := naive(x)
			if index >= 0 {
				fmt.Printf("Majority element is %d", x[index])
			} else {
				fmt.Print("Element not found")
			}
		}
	case "hashing":
		{
			// add +1 to hash if element is encountered
			hashing(x)

		}
	case "boyer":
		// find the majority using boyer moore majority algorithm
		/* if the counter is zero we set m to ekement and couter to one.
		otherwise if the couter is not zero and x[i] is m then increment counter else decrement counter */
		{
			boyer(x)
		}
	}
}

func naive(x []int) int {

	length := len(x)
	for i := 0; i < length/2; i++ {
		count := 1
		for j := i + 1; j < length; j++ {
			if x[i] == x[j] {
				count++
			}
			if count > length/2 {
				return i
			}
		}

	}
	return -1
}

func hashing(x []int) {
	var mapArr = make(map[int]int)
	for i := 0; i < len(x); i++ {
		if value, ok := mapArr[x[i]]; ok {
			mapArr[x[i]] = value + 1
		} else {
			mapArr[x[i]] = 1
		}
	}
	for key, element := range mapArr {
		if element > len(x)/2 {
			fmt.Printf("Majority element is %d", key)
		}
	}
}

func boyer(x []int) {
	m := -1
	i := 0
	for j := 0; j < len(x); j++ {
		if i == 0 {
			m = x[j]
			i = 1
		} else if m == x[j] {
			i++
		} else {
			i--
		}
	}
	fmt.Printf("Element is %d", m)
}
