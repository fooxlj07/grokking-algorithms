package binaryseach

func BinarySearch(arr []int, element int) int {
	var mid int
	start, end := 0, len(arr)-1
	for start < end {
		mid = (end + start) / 2
		if element == arr[mid] {
			return mid
		}
		if element < arr[mid] {
			end = mid - 1
		}
		if element > arr[mid] {
			start = mid + 1
		}
	}
	return -1
}
