package main

// Bubble - Loop through the array check each adjacent number if the next one is bigger than the current then switch them
func Bubble(arr []int) []int {

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				changed = true
			}
		}
	}

	return arr
}

// Insertion - Loop from the second item then from the first if left is bigger than right then swap
func Insertion(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}

	return arr
}

func merge(l, r []int) []int {
	result := make([]int, 0, len(l)+len(r))

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(result, r...)
		}
		if len(r) == 0 {
			return append(result, l...)
		}
		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	return result
}

// Merge - Recursively split the array until all is left are many arrays of 1 item, then stitch them back together in order
func Merge(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	l := Merge(arr[:mid])
	r := Merge(arr[mid:])
	return merge(l, r)
}
