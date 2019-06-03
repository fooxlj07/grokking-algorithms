package main

import (
	"fmt"
	"chapter1"
)

func main() {
	arr := []int{4,1,5,3,6,2,7}
	fmt.Println(chapter1.BinarySearch(arr, 1))
}