package utils

import (
	"fmt"
)

func PassSlice() {
	passSlice1()
	fmt.Println("----------")
	passSlice2()
}

func passSlice1() {
	arr := []byte{1, 2, 3, 4, 5}

	fmt.Printf("%p, %p, %p, %p\n", arr, &arr[0], &arr[1], &arr[2])
	// 0xc0000b0058, 0xc0000b0058, 0xc0000b0059, 0xc0000b005a
	passValue(arr)
	// 0xc0000b0058, 0xc0000b0058, 0xc0000b0059, 0xc0000b005a
	// len=8,cap=16
	fmt.Println(arr)                                  // [1 11 3 4 5]
	fmt.Printf("len=%d,cap=%d\n", len(arr), cap(arr)) // len=5,cap=5

	passPointer(&arr)
	// pointer address: 0xc0000b0058
	// address of pointer: 0xc0000bc030
	// len=8,cap=16
	fmt.Println(arr)                                  // [1 11 3 4 5 77 88 99]
	fmt.Printf("len=%d,cap=%d\n", len(arr), cap(arr)) // len=8,cap=16
}

func passSlice2() {
	arr := make([]byte, 0, 10)
	arr = append(arr, 1, 2, 3, 4, 5)

	fmt.Printf("%p, %p, %p, %p\n", arr, &arr[0], &arr[1], &arr[2])
	// 0xc0000b00b0, 0xc0000b00b0, 0xc0000b00b1, 0xc0000b00b2
	passValue(arr)
	// 0xc0000b00b0, 0xc0000b00b0, 0xc0000b00b1, 0xc0000b00b2
	// len=8,cap=10
	fmt.Println(arr)                                  // [1 11 3 4 5]
	fmt.Printf("len=%d,cap=%d\n", len(arr), cap(arr)) // len=5,cap=10

	passPointer(&arr)
	// pointer address: 0xc0000b00b0
	// address of pointer: 0xc0000bc0c0
	// len=8,cap=10
	fmt.Println(arr)      // [1 11 3 4 5 77 88 99]
	fmt.Printf("len=%d,cap=%d\n", len(arr), cap(arr)) // len=8,cap=10

}

func passValue(val []byte) {
	fmt.Printf("%p, %p, %p, %p\n", val, &val[0], &val[1], &val[2])
	val[1] = 11
	val = append(val, 77, 88, 99)                     // this doesn't affect the original value
	fmt.Printf("len=%d,cap=%d\n", len(val), cap(val))
}

func passPointer(val *[]byte) {
	fmt.Printf("pointer address: %p\n", *val)
	fmt.Printf("address of pointer: %p\n", val)
	(*val)[1] = 11
	*val = append(*val, 77, 88, 99)
	fmt.Printf("len=%d,cap=%d\n", len(*val), cap(*val))
}
