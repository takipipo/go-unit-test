package sum_test

import (
	"go-unit-test/sum"
	"testing"
)

type testCase struct {
	name           string
	numbers        []int
	expectedResult int
}

func TestInts(t *testing.T) {
	testCases := []testCase{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := sum.Ints(tc.numbers...)
			expectedResult := tc.expectedResult
			if actualResult != expectedResult {
				t.Errorf("The expected result of %v is %v but got %v", tc.name, expectedResult, actualResult)
			}
		})
	}
}
