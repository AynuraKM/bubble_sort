package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 2)
	arr := []int{}

	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(50))
	}

	fmt.Println(arr)

	bubbleSort(arr)

	fmt.Println(arr)

	elapsed := time.Since(start)
	fmt.Printf("sort took %s", elapsed)
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
