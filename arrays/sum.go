package arrays

// https://blog.golang.org/go-slices-usage-and-internals
// In Go, there are distinction between Array and Slice:
//
// Array: a fixed capacity
//        > var a [4]int   // {0, 0, 0, 0}
//        > a[0] = 1       // {1, 0, 0, 0}
//
//        > b := [...]int {1, 2}  // ... let compiler fill the size (2)
//
// Slice : a dynamically-resized array
//
//

// Sum takes a slice and returns a sum of all its elements.
func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += array[i]
	// }
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes multiple slices and return an array of their sums.
func SumAll(numbersToSum ...[]int) (sums []int) {
	// size := len(numbersToSum)
	// sums = make([]int, size)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	// return
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails takes multiple slices and return an array of their sums.
func SumAllTails(nums ...[]int) (sums []int) {
	for _, numbers := range nums {
		sums = append(sums, Sum(numbers[1:]))
	}
	return sums
}
