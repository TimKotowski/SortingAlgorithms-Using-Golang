package main

import (
	"fmt"

	"godata/data/bubble"
	"godata/data/selection"
)

func main() {
	array := []int{8, 5, 2, 1, 23, 6, 88, 4, 88, 979, 9, 5, 6, 3}
	fmt.Println("not sorted", array)
	bubble.BubbleSort(array	)
	fmt.Println("bublesort sorted", array)
	arrayTwo := []int{2,4,6,5,3,1}
	fmt.Println("not sorted", arrayTwo)
	selection.SelectionSort(arrayTwo)
	fmt.Println("selectionsort sorted", arrayTwo)
}
