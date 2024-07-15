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
		value: intValue,
	}

	doSwitchType(containsInt)  // default
	doSwitchType(&containsInt) // default

	fmt.Println("--- containsAny.value ---")
	doSwitchType(containsInt.value)    // int
	doSwitchType(&containsInt.value)   // default
	doSwitchType(&(containsInt.value)) // default

	fmt.Printf("%p, %p, %p\n", &containsInt, &containsInt.value, &(containsInt.value)) // 0xc000014180, 0xc000014180, 0xc000014180
	checkStructInfo(containsInt)
	// struct size for [anyInStruct] is 16 (any: 16)
	// value     : (offset: 0, align: 8, size: 16)

	containsPointerInt := anyInStruct{
		value: &intValue,
	}
	containsPointerStruct := anyInStruct{
		value: &yuto,
	}
	doSwitchType(containsPointerInt.value)    // int pointer
	doSwitchType(containsPointerStruct.value) // (person pointer) name: Yuto

	fmt.Println("--- update data ---")
	updateData(&containsInt.value)
	fmt.Println(containsInt.value) // 10
	updateData(&containsStruct.value)
	fmt.Println(containsStruct.value) // {Yuto 0}
	updateData(&containsPointerStruct.value)
	fmt.Println(containsPointerStruct.value) // &{someone2 0}

	fmt.Println("--- update person ---")
	person1 := person{
		Name: "me",
		Age:  55,
	}
	updatePerson(&person1)
	fmt.Println(person1)

	personType, ok := containsStruct.value.(person)
	if !ok {
		fmt.Println("failed to convert")
	} else {
		updatePerson(&personType)
		fmt.Println(personType)
	}

	fmt.Println("--- loop ---")
	intValue2 := 22
	yuto = person{
		Name: "YUTO",
		Age:  36,
	}
	items := []anyInStruct{
		{
			value: 10,
		},
		{
			value: &intValue2,
		},
		{
			value: person{
				Name: "yuto",
			},
		},
		{
			value: &yuto,
		},
	}

	for _, item := range items {
		updateData(&item.value)
		printBySwitchType(item.value)
	}
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

func printBySwitchType(value any) {
	switch val := value.(type) {
	case int:
		fmt.Println(val)
	case *int:
		fmt.Println(*val)
	case person:
		fmt.Println(val)
	case *person:
		fmt.Println(val)
	default:
		fmt.Println("default")
	}
}

func updateData(value *any) {
	switch val := (*value).(type) {
	case int:
		fmt.Print("int  - ")
		val = 111
	case *int:
		fmt.Print("pointer int  - ")
		newValue := 999
		*val = newValue
	case person:
		fmt.Print("person  - ")
		val.Name = "someone"
	case *person:
		fmt.Print("pointer person  - ")
		val.Name = "someone2"
	default:
		fmt.Println("do nothing")
	}
}

func updatePerson(value *person) {
	value.Name = "new name"
	value.Age = 99
	value.Gender = "male"
}
