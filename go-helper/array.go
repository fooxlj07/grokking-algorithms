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

func DeleteFromIntByIndex(arr []int, index int) []int {
	if index >= len(arr) {
		return arr
	}
	if index == 0 {
		return arr[1:]
	}
	if index == len(arr)-1 {
		return arr[:index]
	}
	return append(arr[:index], arr[index+1:]...)
}

func DeleteFromStrByIndex(arr []string, index int) []string {
	if index >= len(arr) {
		return arr
	}
	if index == 0 {
		return arr[1:]
	}
	if index == len(arr)-1 {
		return arr[:index]
	}
	return append(arr[:index], arr[index+1:]...)
}

func UnionStrSlice(arr1 []string, arr2 []string) []string {
	m := map[string]bool{}
	union := []string{}
	for _, a := range arr1 {
		m[a] = true
	}
	for _, b := range arr2 {
		if _, ok := m[b]; !ok {
			union = append(union, b)
		}
	}

	return append(arr1, union...)
}

func IntersectionStrSlice(arr1 []string, arr2 []string) []string {
	m := map[string]bool{}
	intersection := []string{}
	for _, a := range arr1 {
		m[a] = true
	}
	for _, b := range arr2 {
		if _, ok := m[b]; ok {
			intersection = append(intersection, b)
		}
	}
	return intersection
}

func RemoveStrSlice1FromSlice2(arr1 []string, arr2 []string) []string {
	m := map[string]bool{}
	arr := []string{}
	for _, a := range arr2 {
		m[a] = true
	}
	for _, b := range arr1 {
		if !m[b] {
			arr = append(arr, b)
		}
	}
	return arr
}
