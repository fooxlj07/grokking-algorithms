package chapter4

func Sum(arr []int) int{
	if len(arr) == 0 {
		return 0
	}
	if len(arr) ==1{
		return arr[0]
	}
	return arr[0] + Sum(arr[1:])
}

func Max(arr []int) int{
	if len(arr) == 0{
		return 0
	}
	if len(arr) == 1{
		return arr[0]
	}
	if arr[0] > Max(arr[1:]){
		return arr[0]
	}
	return Max(arr[1:])
}

func BinarySearch(arr []int, element, start, end int) int{
	if start >= end {
		return -1
	}
	mid := end + start / 2
	if arr[mid] == element{
		return mid
	}
	if element > arr[mid] {
		return BinarySearch(arr, element, mid+1, end)
	}
	return BinarySearch(arr, element , start, mid-1)
}

func QuickSort(arr []int) []int{
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left, right := []int{}, []int{}
	for _, a := range arr[1:] {
		if a <= pivot {
			left = append(left, a)
		}else{
			right = append(right, a)
		}
	}
	result := append(QuickSort(left), pivot)
	return append(result, QuickSort(right)...)
}