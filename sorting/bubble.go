package sorting

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
