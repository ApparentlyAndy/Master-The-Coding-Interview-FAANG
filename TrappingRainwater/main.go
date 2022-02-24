package main

import "fmt"

// func findLargest(arr []int) int {
// 	largest := 0
// 	for _, num := range arr {
// 		if num > largest {
// 			largest = num
// 		}
// 	}
// 	return largest
// }

// func trappingRainwater(height []int) int {
// 	totalRain := 0
// 	largest := findLargest(height)

// 	for i := 1; i <= largest; i++ {
// 		currentLevel := []int{}
// 		for j := 0; j < len(height); j++ {
// 			if height[j] >= i {
// 				currentLevel = append(currentLevel, 1)
// 				continue
// 			}
// 			currentLevel = append(currentLevel, 0)
// 		}

// 		for k := 0; k < len(currentLevel); k++ {
// 			if currentLevel[k] != 0 {
// 				for m := k + 1; m < len(currentLevel); m++ {
// 					if currentLevel[m] >= currentLevel[k] {
// 						totalRain += (m - k - 1)
// 						k = m
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return totalRain
// }

func trappingRainwater(heights []int) int {
	left := 0
	right := len(heights) - 1
	leftMax := 0
	rightMax := 0
	total := 0

	for left < right {
		if heights[left] < heights[right] {
			if heights[left] >= leftMax {
				leftMax = heights[left]
			} else {
				total += leftMax - heights[left]
			}
			left++
		} else {
			if heights[right] >= rightMax {
				rightMax = heights[right]
			} else {
				total += rightMax - heights[right]
			}
			right--
		}
	}
	return total
}

func main() {
	fmt.Println(trappingRainwater([]int{4, 2, 0, 3, 2, 5}))
}
