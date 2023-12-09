package utils

import (
	"fmt"
	"reflect"
)

func SliceTest() {
	var list []int

	fmt.Println(list) // []
	list = append(list, 9)
	fmt.Println(list) // [9]
	list = append(list, 8, 7, 6)
	fmt.Println(list) // [9 8 7 6]

	addNineByValue(list)
	fmt.Println(list) // [9 8 7 6]
	addNineByPointer(&list)
	fmt.Println(list) // [9 8 7 6 9]

	list = nil
	fmt.Println(list)      // []
	fmt.Println(len(list)) // 0

	var emptyList []int
	list = append(list, emptyList...)
	fmt.Println(list) // []

	mySlice := []int{9, 8, 7, 6}
	list = append(list, mySlice...)
	fmt.Println(list) // [9 8 7 6]

	list = nil

	myArray := [4]int{5, 4, 3, 2}
	list = append(list, myArray[:]...)
	fmt.Println(list) // [5 4 3 2]

	fmt.Println(reflect.TypeOf(myArray).Kind())    // array
	fmt.Println(reflect.TypeOf(myArray[:]).Kind()) // slice

	list = nil

	fmt.Println("--- COLON ---")

	list = append(list, 1, 2, 3, 4)
	fmt.Println(list)      // [1 2 3 4]
	fmt.Println(list[:])   // [1 2 3 4]
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

	fmt.Println("--- Some --- ")
	fmt.Println(Some(list, func(value int) bool { return value < 2 }))  // true
	fmt.Println(Some(list, func(value int) bool { return value > 4 }))  // false
	fmt.Println(Some(list, func(value int) bool { return value == 4 })) // true

	fmt.Println("--- Contains --- ")
	fmt.Println(Contains(list, 2)) // true
	fmt.Println(Contains(list, 4)) // true
	fmt.Println(Contains(list, 5)) // false

	fmt.Println("--- Find --- ")
	fmt.Println(*Find(list, func(value int) bool { return value < 2 }))  // 1
	fmt.Println(Find(list, func(value int) bool { return value > 4 }))   // <nil>
	fmt.Println(*Find(list, func(value int) bool { return value == 4 })) // 4

	fmt.Println("--- Find2 --- ")
	value, ok := Find2(list, func(value int) bool { return value < 2 })
	fmt.Printf("value: %v, ok: %t\n", value, ok)
	value, ok = Find2(list, func(value int) bool { return value == 5 })
	fmt.Printf("value: %v, ok: %t\n", value, ok)

	map1 := map[string]int{"first": 1, "second": 2}
	map2 := map[string]int{"three": 3, "four": 4}
	valueMap, ok := Find2([]map[string]int{map1, map2}, func(valueMap map[string]int) bool {
		_, ok := valueMap["first"]
		return ok
	})
	fmt.Printf("value: %v, ok: %t\n", valueMap, ok) // value: map[first:1 second:2], ok: true

	valueMap, ok = Find2([]map[string]int{map1, map2}, func(valueMap map[string]int) bool {
		_, ok := valueMap["notExist"]
		return ok
	})
	fmt.Printf("value: %v, ok: %t\n", valueMap, ok) // value: map[], ok: false
}

func addNineByValue(array []int) {
	array = append(array, 9)
}
func addNineByPointer(array *[]int) {
	*array = append(*array, 9)
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

func Some[T any](array []T, callback func(T) bool) bool {
	for _, value := range array {
		if callback(value) {
			return true
		}
	}

	return false
}

func Contains[T comparable](array []T, searchValue T) bool {
	for _, value := range array {
		if value == searchValue {
			return true
		}
	}

	return false
}

func Find[T any](array []T, callback func(T) bool) *T {
	for _, value := range array {
		if callback(value) {
			return &value
		}
	}

	return nil
}

func Find2[T any](array []T, callback func(T) bool) (T, bool) {
	for _, value := range array {
		if callback(value) {
			return value, true
		}
	}

	var zeroValue T
	return zeroValue, false
}

func RunSliceArrayDeclaration() {
	slice1 := []int{}
	fmt.Printf("nil == %t; len:cap = %d:%d\n", slice1 == nil, len(slice1), cap(slice1))

	var slice2 []int
	fmt.Printf("nil == %t; len:cap = %d:%d\n", slice2 == nil, len(slice2), cap(slice2))

	slice3 := make([]int, 0)
	fmt.Printf("nil == %t; len:cap = %d:%d\n", slice3 == nil, len(slice3), cap(slice3))

	array1 := [100]int{}
	fmt.Printf("len:cap = %d:%d\n", len(array1), cap(array1))

	array2 := make([]int, 0, 100)
	fmt.Printf("len:cap = %d:%d\n", len(array2), cap(array2))

	array3 := make([]int, 100, 100)
	fmt.Printf("len:cap = %d:%d\n", len(array3), cap(array3))
}
