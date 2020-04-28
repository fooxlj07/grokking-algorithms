package helper

func Max(arr []int) (int, int) {
	max, index := 0, -1
	for i, num := range arr {
		if num >= max {
			max = num
			index = i
		}
	}
	return index, max
}

func Min(arr []int) (int, int) {
	min, index := arr[0], -1
	for i, num := range arr {
		if num <= min {
			min = num
			index = i
		}
	}
	return index, min
}

func DeleteByIndex(arr []int, index int) []int {
	if index >= len(arr) || len(arr) <= 1 {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}
