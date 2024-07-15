package utils

import "fmt"

func RunSwitchType() {
	intValue := int(10)
	doSwitchType(intValue)  // int
	doSwitchType(&intValue) // int pointer

	fmt.Println("--- person ---")
	yuto := person{
		Name: "Yuto",
	}
	doSwitchType(yuto)  // person
	doSwitchType(&yuto) // person pointer

	type anyInStruct struct {
		value any
	}

	fmt.Println("--- struct is assigned to any---")
	containsStruct := anyInStruct{
		value: yuto,
	}

	doSwitchType(containsStruct)  // default
	doSwitchType(&containsStruct) // default

	fmt.Println("--- containsAny.value ---")
	doSwitchType(containsStruct.value)    // person
	doSwitchType(&containsStruct.value)   // default
	doSwitchType(&(containsStruct.value)) // default

	fmt.Println("--- int is assigned to any---")
	containsInt := anyInStruct{
		value: int(10),
	}

	doSwitchType(containsInt)  // default
	doSwitchType(&containsInt) // default

	fmt.Println("--- containsAny.value ---")
	doSwitchType(containsInt.value)    // int
	doSwitchType(&containsInt.value)   // default
	doSwitchType(&(containsInt.value)) // default

	// fmt.Println(containsAny.value)
	// fmt.Println(&containsAny.value)

}

func doSwitchType(value any) {
	switch val := value.(type) {
	case int:
		fmt.Println("int")
	case *int:
		fmt.Println("int pointer")
	case person:
		fmt.Printf("(person) name: %s\n", val.Name)
	case *person:
		fmt.Printf("(person pointer) name: %s\n", val.Name)
	default:
		fmt.Println("default")
	}
}
