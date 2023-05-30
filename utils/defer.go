package utils

import "fmt"

func RunDefer() {
	fmt.Println("--- in for loop ---")
	for i := 0; i < 5; i++ {
		fmt.Printf("---- start(%d)\n", i)
		defer func(index int) {
			fmt.Println(index)
		}(i)
		fmt.Printf("---- end(%d)\n", i)
	}

	fmt.Println("--- in ---")
	for i := 0; i < 5; i++ {
		fmt.Printf("---- start(%d)\n", i)
		cb := doIt(i)
		defer cb(i)
		fmt.Printf("---- end(%d)\n", i)
	}

	{
		fmt.Println("starts")
		defer fmt.Println("hello")
		fmt.Println("ends")
	}
	fmt.Println("exit...")

}

func doIt(index int) func(index int) {
	fmt.Printf("in doIt(%d)", index)
	return func(index int) { fmt.Printf("internal(%d)", index) }
}
