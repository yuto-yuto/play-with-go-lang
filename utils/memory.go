package utils

import (
	"fmt"
	"reflect"
	"unsafe"

	"golang.org/x/exp/constraints"
)

type empty struct{}

type memory1[T constraints.Ordered] struct {
	boolValue   bool
	middleValue T
	boolValue2  bool
}

type memory2[T constraints.Ordered] struct {
	boolValue  bool
	boolValue2 bool
	lastValue  T
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
	memory1 memory1[int32]
	memory2 memory2[int32]
}

type hasMapAndSlice struct {
	boolValue  bool
	mapValue   map[int8]int64
	sliceValue []int16
}

const (
	flagBool1 bool = false
	flagInt1  int  = 1
)

func checkStructInfo(obj any) {
	value := reflect.ValueOf(obj)
	fmt.Printf("struct size for [%s] is %d (any: %d)\n", value.Type().Name(), value.Type().Size(), unsafe.Sizeof(obj))

	for i := 0; i < value.Type().NumField(); i++ {
		typeField := value.Type().Field(i)
		field := value.Field(i)
		fmt.Printf("  %-10s: (offset: %d, align: %d, size: %d)\n", typeField.Name, typeField.Offset, field.Type().Align(), field.Type().Size())

	}
}

func RunMemory() {
	fmt.Println(unsafe.Sizeof(memory1[int8]{}))  // 3
	fmt.Println(unsafe.Sizeof(memory1[int16]{})) // 6
	fmt.Println(unsafe.Sizeof(memory1[int32]{})) // 12
	fmt.Println(unsafe.Sizeof(memory1[int64]{})) // 24

	fmt.Println()

	fmt.Println(unsafe.Sizeof(memory2[int8]{}))  // 3
	fmt.Println(unsafe.Sizeof(memory2[int16]{})) // 4
	fmt.Println(unsafe.Sizeof(memory2[int32]{})) // 8
	fmt.Println(unsafe.Sizeof(memory2[int64]{})) // 16

	fmt.Println()

	fmt.Println(unsafe.Sizeof(hasStruct{})) // 20
	checkStructInfo(hasMapAndSlice{})

	fmt.Println()
	checkStructInfo(memory1[int8]{})
	checkStructInfo(memory1[int16]{})
	checkStructInfo(memory1[int32]{})
	checkStructInfo(memory1[int64]{})

	fmt.Println()
	checkStructInfo(memory2[int8]{})
	checkStructInfo(memory2[int16]{})
	checkStructInfo(memory2[int32]{})
	checkStructInfo(memory2[int64]{})

	fmt.Println()
	checkStructInfo(memory1[string]{})
	checkStructInfo(memory2[string]{})

	fmt.Println()
	strObj := memory1[string]{}
	fmt.Println()
	fmt.Printf("%p, size: %d, align: %d\n", &strObj, unsafe.Sizeof(strObj), unsafe.Alignof(strObj))
	checkStructInfo(strObj)
	strObj.middleValue = "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
	checkStructInfo(strObj)
	fmt.Printf("%p, size: %d, align: %d\n", &strObj, unsafe.Sizeof(strObj), unsafe.Alignof(strObj))

	fmt.Println()

	arr := make([]int64, 0, 3)
	fmt.Printf("%p, size: %d, align: %d\n", &arr, unsafe.Sizeof(arr), unsafe.Alignof(arr))
	arr = append(arr, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	arr = append(arr, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Printf("%p, size: %d, align: %d\n", &arr, unsafe.Sizeof(arr), unsafe.Alignof(arr))
}
