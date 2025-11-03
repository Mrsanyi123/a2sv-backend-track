// Task: Sum of Numbers
// Write a Go function that takes a slice of integers as input and returns the sum of all the numbers. If the slice is empty, the function should return 0.
// [Optional]: Write a test for your function

package main

func SumOfNumbers(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}