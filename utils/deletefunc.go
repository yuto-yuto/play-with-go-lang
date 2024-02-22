package utils

import (
	"fmt"
	"slices"
)

func runDeleteFunc1() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array) // [1 2 3 4 5]
	deleted := slices.DeleteFunc(array, func(i int) bool { return i == 2 || i == 3 })
	fmt.Println(array)   // [1 4 5 4 5]
	fmt.Println(deleted) // [1 4 5]
}

func runDeleteFunc2() {
	show := func(data []*int) {
		fmt.Print("[")
		for _, value := range data {
			fmt.Printf("%d ", *value)
		}
		fmt.Println("]")
	}
	one := 1
	two := 2
	three := 3
	four := 4
	five := 5
	array := []*int{&one, &two, &three, &four, &five}
	fmt.Println(array) // [0xc0000ba0e8 0xc0000ba0f0 0xc0000ba0f8 0xc0000ba100 0xc0000ba108]
	show(array)        // [1 2 3 4 5 ]
	deleted := slices.DeleteFunc(array, func(i *int) bool { return *i == 2 || *i == 3 })
	fmt.Println(array)   // [0xc0000ba0e8 0xc0000ba100 0xc0000ba108 0xc0000ba100 0xc0000ba108]
	show(array)          // [1 4 5 4 5 ]
	fmt.Println(deleted) // [0xc0000ba0e8 0xc0000ba100 0xc0000ba108]
	show(deleted)        // [1 4 5 ]

	fmt.Printf("%p, %p\n", array[3], array[4])   // 0xc0000ba100, 0xc0000ba108
	fmt.Printf("%d, %d\n", *array[3], *array[4]) // 4, 5
	clear(array[len(deleted):])
	fmt.Printf("%d, %d\n", four, five)         // 4, 5
	fmt.Printf("%d, %d\n", array[3], array[4]) // 0, 0
}

func RunDeleteFunc() {
	runDeleteFunc1()
	runDeleteFunc2()
}
