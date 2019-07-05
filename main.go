package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{}

	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(50))
	}

	fmt.Println(arr)

	bubbleSort(arr)

	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for i := 1; i < len(arr)-j; i++ {
				if arr[i-1] < arr[i] {
					arr[i-1], arr[i] = arr[i], arr[i-1]
				}
			}
		}
	}
}
