package chapter1

func BinarySearch(arr []int, element int) int{
	var mid int
	start, end := 0, len(arr)-1
	for ; mid > 0 && mid < len(arr){
		mid = (end - start) /2
		if element == arr[mid] {
			return mid
		}
		if element < arr[mid]{
			end = mid
		}
		if element > arr[mid]{
			start = mid
		}
	}
	return -1
}