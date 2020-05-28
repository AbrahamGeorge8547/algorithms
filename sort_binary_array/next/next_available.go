package next 


import "fmt"

func Sort(arr []int) {
	k:=0
	for i:=0; i< len(arr); i++ {
		if arr[i] == 0 {
			arr[k]=0
			k++
		}
	}
	for j:=k; j<len(arr); j++ {
		arr[j]=1
	}
	fmt.Println(arr)
}