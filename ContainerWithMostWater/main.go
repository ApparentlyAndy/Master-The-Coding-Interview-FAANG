package main

import "fmt"

// Test cases [7, 1, 2, 3, 9] = 4 x 7 = 28
// Test cases [] = 0
// Test cases [7] = 0
// Test cases [6, 9, 3, 4, 5, 8] = 8 x 4 = 32

func findMinimum(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func containerWithMostWater(array []int) int {
	maxArea := 0
	start := 0
	end := len(array) - 1
	if len(array) <= 1 {
		return maxArea
	}

	for (end - start) > 1 {
		min := findMinimum(array[start], array[end])
		width := end - start
		result := min * width
		if result > maxArea {
			maxArea = result
		}

		if min == array[start] {
			start++
		} else {
			end--
		}
	}
	return maxArea

}

func main() {
	fmt.Println(containerWithMostWater([]int{}))
}
