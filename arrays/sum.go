package main

// Sum returns the sum of any size of a collection
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}
