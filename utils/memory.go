package utils

import (
	"fmt"
	"unsafe"
)

type memory1 struct {
	boolValue  bool
	int32Value int32
	boolValue2 bool
}
type memory1_2 struct {
	boolValue  bool
	int32Value int64
	boolValue2 bool
}

type memory2 struct {
	boolValue  bool
	boolValue2 bool
	int32Value int32
}
type memory2_2 struct {
	boolValue  bool
	boolValue2 bool
	int32Value int64
}

type memory3 struct {
	boolValue  bool
	boolValue2 bool
	boolValue3 bool
	boolValue4 bool
	// boolValue5 bool
	// boolValue6 bool
	// boolValue7 bool
	// boolValue8 bool
	int32Value int32
}

type hasStruct struct {
	memory1 memory1
	memory2 memory2
}

const (
	flagBool1 bool = false
	flagInt1  int  = 1
)

func RunMemory() {
	fmt.Println(unsafe.Sizeof(memory1{}))   // 12
	fmt.Println(unsafe.Sizeof(memory1_2{})) // 24
	fmt.Println(unsafe.Sizeof(memory2{}))   // 8
	fmt.Println(unsafe.Sizeof(memory2_2{})) // 16
	fmt.Println(unsafe.Sizeof(memory3{}))   // 8
	fmt.Println(unsafe.Sizeof(hasStruct{})) // 20

	fmt.Println()

	obj := memory1{}
	fmt.Printf("%p\n", &obj)            // 0xc000180000
	fmt.Printf("%p\n", &obj.boolValue)  // 0xc000180000
	fmt.Printf("%p\n", &obj.int32Value) // 0xc000180004
	fmt.Printf("%p\n", &obj.boolValue2) // 0xc000180008

	fmt.Println()

	obj2 := memory2{}
	fmt.Printf("%p\n", &obj2)            // 0xc000180020
	fmt.Printf("%p\n", &obj2.boolValue)  // 0xc000180020
	fmt.Printf("%p\n", &obj2.int32Value) // 0xc000180024
	fmt.Printf("%p\n", &obj2.boolValue2) // 0xc000180021

	fmt.Println()

	fmt.Println(unsafe.Sizeof(flagBool1))
	fmt.Println(unsafe.Sizeof(flagInt1))
}
