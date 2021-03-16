package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 2, 5, 7, -2, 15, 14, -15},
			expected: 15,
		},
		{
			numbers:  []int{0, -2, -10},
			expected: 0,
		},
		{
			numbers:  []int{},
			expected: 0,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := Max(testCase.numbers)

		t.Logf("Calling Max(%v), result %d\n",
			testCase.numbers, result)

		// Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %d, got %d",
				testCase.expected, result)
		}
	}

}

func TestMaxIndex(t *testing.T) {
	// Arrange
	numbers := []int{1, 2, 5, 7, -2, 15, 14, -15}
	expected := 5

	//Act
	result := MaxIndex(numbers)

	// Assert
	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d",
			expected, result)
	}
}
