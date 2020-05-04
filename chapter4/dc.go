package chapter4

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

func Run() {
	fmt.Println(divideTheGround(24, 640))
	fmt.Println(sum([]int{2, 4, 6}))
}
