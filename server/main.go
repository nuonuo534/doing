package main

import (
	"fmt"
	"server/algorithm"

	_ "server/setting"
)

func main() {
	arr := []int{6, 7, 2, 34, 5, 8, 9, 1}
	algorithm.BubbleSort(arr)
	fmt.Printf("int: %v\n", algorithm.Rescuive(3))
	fmt.Printf("arr : %v\n", arr)
}
