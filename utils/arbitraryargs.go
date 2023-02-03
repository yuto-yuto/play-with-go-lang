package utils

import (
	"fmt"
	"reflect"
)

func arbitraryNumbers(args ...int) {
	fmt.Println("--- arbitraryNumbers ---")
	fmt.Printf("0: %d, 1: %d\n", args[0], args[1])
	for index, arg := range args {
		fmt.Printf("%d: %d, ", index, arg)
	}
	fmt.Println()
}
func arbitraryStrings(args ...string) {
	fmt.Println("--- arbitraryStrings ---")
	fmt.Printf("0: %s, 1: %s\n", args[0], args[1])
	for index, arg := range args {
		fmt.Printf("%d: %s, ", index, arg)
	}
	fmt.Println()
}
func arbitraryAny(args ...any) {
	fmt.Println("--- arbitraryAny ---")
	fmt.Printf("0: %v, 1: %v\n", args[0], args[1])
	for index, arg := range args {
		fmt.Printf("%d: %v, ", index, arg)
	}
	fmt.Println()
}

func arbitraryObj(args ...any) {
	for index, obj := range args {
		fmt.Printf("%d: %#v\n", index, obj)
		if reflect.TypeOf(obj) == reflect.TypeOf(unknownObj1{}) {
			fmt.Printf("Name: %s, ", getValueOf(obj, "Name"))
			fmt.Printf("Prop1: %s\n", getValueOf(obj, "Prop1"))
		} else if reflect.TypeOf(obj) == reflect.TypeOf(unknownObj2{}) {
			castedObj := obj.(unknownObj2)
			fmt.Printf("Name: %s, Prop1: %s, Count: %d\n", castedObj.Obj.Name, castedObj.Obj.Prop1, castedObj.Count)
		}
	}
}

func RunArbitraryNumberOfArgs() {
	arbitraryNumbers(22, 77, 44, 57, 92, 23)
	arbitraryStrings("Hello", "Wor", "ld", "!!!")
	arbitraryAny("Hello", 123, "Wor", 987, "ld", 888, "!!!")
	obj1 := unknownObj1{
		Name: "Yuto", Prop1: "Hello World!",
	}
	obj2 := unknownObj2{
		Obj: unknownObj1{
			Name: "Agent", Prop1: "ABCDE",
		},
		Count: 66,
	}
	arbitraryObj(obj1, obj2)
}
