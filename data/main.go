package main

import (
	"fmt"
	"godata/data/bubble"
)

func main() {
	array := []int{8, 5, 2, 1, 23, 6, 88, 4, 88, 979, 9, 5, 6, 3, 21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	fmt.Println("not sorted", array)
	bubble.BubbleSort(array)
	fmt.Println("sortred", array)
}
