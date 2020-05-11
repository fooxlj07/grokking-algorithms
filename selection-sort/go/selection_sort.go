package selectionsort

import (
	helper "grokking-algorithms/go-helper"
)

func SelectionSort(arr []int) []int {
	sortedArr := []int{}
	l := len(arr)
	for i := 0; i < l; i++ {
		maxIndex, max := helper.Max(arr)
		sortedArr = append(sortedArr, max)
		arr = helper.DeleteByIndex(arr, maxIndex)
	}
	return sortedArr
}
