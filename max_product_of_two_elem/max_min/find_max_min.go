package max_min

import "fmt"

func Find(x []int) {
	max_1, max_2 := x[0], -^int(0)

	min_1, min_2 := x[0], ^int(0)
	for i := 1; i < len(x); i++ {
		if x[i] > max_1 {
			max_2 = max_1
			max_1 = x[i]
		} else if x[i] > max_2 {
			max_2 = x[i]
		} else if x[i] < min_1 {
			min_2 = min_1
			min_1 = x[i]
		} else if x[i] < min_2 {
			min_2 = x[i]
		}
	}
	if (max_1 * max_2) > (min_1 * min_2) {
		fmt.Printf("max elements are %d and %d", max_1, max_2)
	} else {
		fmt.Printf("max elements are %d and %d", min_1, min_2)
	}
}
