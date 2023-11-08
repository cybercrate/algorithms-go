package search

func Linear(data []int, target int) (int, error) {
	for i, item := range data {
		if item == target {
			return i, nil
		}
	}
	return -1, ErrNotFound
}
