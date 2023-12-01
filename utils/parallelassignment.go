package utils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func RunMultiAssign() {
	array1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v -> ", array1)
	ReverseArray1(array1)
	fmt.Println(array1)
}

func ReverseArray1[T constraints.Ordered](array []T) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func ReverseArray2[T constraints.Ordered](array []T) {
	length := len(array)
	result := make([]T, length, length)
	for index, value := range array {
		result[length-index] = value
	}
}

func generateIntArrays() ([]int, []int) {
	return []int{0, 1, 2, 3, 4}, []int{5, 6, 7, 8, 9}
}
