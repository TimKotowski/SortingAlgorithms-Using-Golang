package bubble

func BubbleSort(array []int) {
	var N int = len(array)
	for i := 0; i < N; i++ {
		if !swap(array) {
			return
		}
	}
}

func swap(array []int) bool {
	var N int = len(array)
	firstIndex := 0
	secondIndex := 1
	didSwap := false
	for secondIndex < N {
		firstNumber := array[firstIndex]
		secondNumber := array[secondIndex]

		if firstNumber > secondNumber {
			array[firstIndex] = secondNumber
			array[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
