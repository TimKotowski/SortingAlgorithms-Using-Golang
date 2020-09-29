package insertion

func InsertionSort(array []int) []int {
	for i := range array {
		j := i
		for j > 0 && array[j] < array[j-1] {
			array[j], array[j-1] = array[j-1], array[j]
			j--
		}
	}
	return array
}
