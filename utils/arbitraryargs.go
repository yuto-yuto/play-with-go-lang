package utils

import "fmt"

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

func RunArbitraryNumberOfArgs() {
	arbitraryNumbers(22, 77, 44, 57, 92, 23)
	arbitraryStrings("Hello", "Wor", "ld", "!!!")
	arbitraryAny("Hello", 123, "Wor", 987, "ld", 888, "!!!")
}
