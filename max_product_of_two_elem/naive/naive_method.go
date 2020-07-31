package naive

import "fmt"

func Find(x []int) {
	maxint := -^int(0)
	maxI, maxJ := -1, -1
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if maxint < x[i]*x[j] {
				maxint = x[i] * x[j]
				maxI, maxJ = i, j
			}
		}
	}
	fmt.Printf("%d%d", maxI, maxJ)

}
