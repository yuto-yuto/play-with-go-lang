package utils

import (
	"fmt"
	"reflect"
)

func runIntParams(args ...int) {
	if args == nil {
		fmt.Println("no args")
	}
	for i := 0; i < len(args); i++ {
		fmt.Printf("%d, ", args[i])
	}
	fmt.Println()
}

func runFuncParams(args ...func(param int) string) {
	if args == nil {
		fmt.Println("no args")
	}

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i](i))
	}
}

func runFuncMixedType(args ...interface{}) {
	if args == nil {
		fmt.Println("no args")
	}

	for i := 0; i < len(args); i++ {
		fmt.Printf("type: %s,\t value: %v\n", reflect.TypeOf(args[i]), args[i])
	}
}

func RunOptionalParams() {
	runIntParams()
	runIntParams(9)
	runIntParams(9, 8, 7)

	fmt.Println()
	runFuncParams(
		func(param int) string {
			return fmt.Sprintf("Hello World%d", param)
		},
		func(param int) string {
			return fmt.Sprintf("Foooooo%d", param)
		},
	)
	fmt.Println()
	runFuncMixedType(true, 12, "hoge")
}
