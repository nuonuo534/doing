package main

import (
	"fmt"
	"server/algorithm"

	_ "github.com/nuonuo534/setting"
)

func main() {
	arr := []int{6, 7, 2, 34, 5, 8, 9, 1}
	algorithm.ArrayBubble(arr)
	fmt.Printf("arr : %v\n", arr)
}