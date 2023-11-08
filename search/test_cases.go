package search

type searchTest struct {
	data          []int
	key           int
	expected      int
	expectedError error
	name          string
}

var sanity = "Sanity"
var absent = "Absent"
var empty = "Empty"

var searchTests = []searchTest{
	// Sanity.
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, nil, sanity},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil, sanity},

	// Absent.
	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, ErrNotFound, absent},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrNotFound, absent},

	// Empty slice.
	{[]int{}, 2, -1, ErrNotFound, empty},
}

func generateBenchmarkTestCase() []int {
	var testCase []int

	for i := 0; i < 1000; i++ {
		testCase = append(testCase, i)
	}
	return testCase
}
