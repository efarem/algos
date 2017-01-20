package sorting

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
