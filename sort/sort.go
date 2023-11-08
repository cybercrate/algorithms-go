package sort

import "github.com/wingmann/algorithms/constraints"

func Bubble[T constraints.Ordered](data []T) []T {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(data)-1; i++ {
			if data[i+1] < data[i] {
				data[i+1], data[i] = data[i], data[i+1]
				swapped = true
			}
		}
	}
	return data
}
