package search

import (
	"errors"
	"testing"
)

func TestLinear(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := Linear(test.data, test.key)

		if actualValue != test.expected {
			t.Errorf(
				"test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'",
				test.name,
				test.data,
				test.key,
				test.expected,
				actualValue)
		}

		if !errors.Is(actualError, test.expectedError) {
			t.Errorf(
				"test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'",
				test.name,
				test.data,
				test.key,
				test.expectedError,
				actualError)
		}
	}
}

func BenchmarkLinear(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Linear(testCase, i)
	}
}
