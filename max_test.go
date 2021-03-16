package main

import "testing"

func TestMax(t *testing.T) {
	// Arrange
	numbers := []int{1, 2, 5, 7, -2, 15, 14, -15}
	expected := 15

	//Act
	result := Max(numbers)

	// Assert
	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d",
			expected, result)
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
