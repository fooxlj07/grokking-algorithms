package dc

import "fmt"

// divide the area into squares which have the same area (as big as possible)
// 1. baseline condition:  one side is another side's multiple
func divideTheGround(width int, height int) int {
	fmt.Println(width, height)
	if height == width {
		return height
	}
	if height <= 0 || width <= 0 {
		return 0
	}
	// make height always bigger than width
	if height < width {
		width, height = height, width
	}
	if height%width == 0 {
		return width
	}

	return divideTheGround(height%width, width)
}

func sum(arr []int) int {
	fmt.Println(arr)
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	return sum(arr[1:]) + arr[0]
}

func max(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	secondMax := max(arr[1:])
	if arr[0] > secondMax {
		return arr[0]
	} else {
		return secondMax
	}
}

// baseline condition: length of array < 2
// split the slice to sub slice, choose a pivot,
// subarry1: less than pivot or equal,  subarry2: greater than pivot
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	less, greater := []int{}, []int{}
	for _, a := range arr[1:] {
		if a <= arr[0] {
			less = append(less, a)
		} else {
			greater = append(greater, a)
		}
	}
	return append(append(quickSort(less), arr[0]), quickSort(greater)...)
}

func Run() {
	fmt.Println(divideTheGround(24, 640))
	fmt.Println(sum([]int{2, 4, 6}))
	fmt.Println(max([]int{2, 7, 6}))
	fmt.Println(quickSort([]int{2, 7, 6, 3, 1}))
}
