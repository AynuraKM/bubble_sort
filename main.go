package main

import "fmt"

func main() {
	var arr1 []int = []int{1,5,2,4,8,6,7,2,3,55}

	for j := 0; j < len(arr1); j++ {
			for i := 1; i < len(arr1)-j; i++ {
				if arr1[i-1] < arr1[i] {
					arr1[i-1], arr1[i] = arr1[i], arr1[i-1]
				}
		}
	}
	fmt.Println(arr1)
}
