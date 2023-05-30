package utils

import "fmt"

func SliceTest() {
	var list []int

	fmt.Println(list)
	var emptyList []int
	list = append(list, emptyList...)
	fmt.Println(list)

	list = append(list, []int{1, 2, 3, 4}...)
	fmt.Println(list)      // [1 2 3 4]
	fmt.Println(list[0:])  // [1 2 3 4]
	fmt.Println(list[0:1]) // [1]
	fmt.Println(list[0:2]) // [1 2]
	fmt.Println(list[2:3]) // [3]
	fmt.Println(list[2:2]) // []
	fmt.Println(list[2:4]) // [3 4]
	fmt.Println(list[:3])  // [1 2 3]

	// fmt.Println(list[4:2]) invalid slice indices: 2 < 4
	// fmt.Println(list[:-3]) invalid argument: index -3 (constant of type int) must not be negative

	fmt.Println()

	fmt.Println("--- Filter --- ")
	// [3 4]
	fmt.Println(Filter(list, func(value int) bool { return value > 2 }))
	// []
	fmt.Println(Filter(list, func(value int) bool { return value > 4 }))

	fmt.Println("--- Select --- ")
	// [1 4 9 16]
	fmt.Println(Select(list, func(value int) int { return value * value }))
	// [result: 1 result: 4 result: 9 result: 16]
	fmt.Println(Select(list, func(value int) string { return fmt.Sprintf("result: %d", value*value) }))

	fmt.Println("--- Contains --- ")
	fmt.Println(Contains(list, func(value int) bool { return value < 2 }))  // true
	fmt.Println(Contains(list, func(value int) bool { return value > 4 }))  // false
	fmt.Println(Contains(list, func(value int) bool { return value == 4 })) // true

}

func Filter[T any](array []T, callback func(T) bool) []T {
	result := []T{}

	for _, value := range array {
		if callback(value) {
			result = append(result, value)
		}
	}

	return result
}

func Select[T any, K any](array []T, callback func(T) K) []K {
	result := make([]K, len(array))

	for index, value := range array {
		result[index] = callback(value)
	}

	return result
}

func Contains[T any](array []T, callback func(T) bool) bool {
	for _, value := range array {
		if callback(value) {
			return true
		}
	}

	return false
}
