package hashing

func FindPair(arr []int, sum int) {

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
