package utils

import "fmt"

type myStruct struct {
	name string
	age  uint64
}

func (m myStruct) String() string {
	return fmt.Sprintf("{ name: %s, age: %d }", m.name, m.age)
}

func StringHandling() {
	str := string("Hello world!")
	fmt.Println(&str) // 0xc000014300
	fmt.Println(str)  // Hello world!
	fmt.Println()

	var strPointer *string
	fmt.Println(strPointer == nil) // true
	temp := string("Hello world!")
	strPointer = &temp
	fmt.Println(strPointer)  // 0xc000014320
	fmt.Println(*strPointer) // Hello world!
	fmt.Println()

	fmt.Printf("&str == strPointer => %t\n", &str == strPointer) // &str == strPointer => false
	fmt.Printf("str == *strPointer => %t\n", str == *strPointer) // str == *strPointer => true
	fmt.Println()

	intValue := 99
	intStr1 := string(intValue)
	fmt.Printf("intStr1: %s\n", intStr1) // c

	intStr2 := fmt.Sprint(intValue)
	fmt.Printf("intStr2: %s\n", intStr2) // 99

	floatValue := 12.345
	floatStr := fmt.Sprint(floatValue)
	fmt.Printf("floatStr: %s\n", floatStr) // 12.345
	fmt.Println()

	arrayValue := [5]int{1, 2, 3, 4, 5}
	arrayStr := fmt.Sprint(arrayValue)
	fmt.Printf("arrayValue: %s\n", arrayStr)
	fmt.Println()

	myStruct := myStruct{name: "Yuto", age: 35}
	fmt.Printf("%%v: %v\n", myStruct)   // %v: {Yuto 35}
	fmt.Printf("%%T: %T\n", myStruct)   // %T: utils.myStruct
	fmt.Printf("%%#v: %#v\n", myStruct) // %#v: utils.myStruct{name:"Yuto", age:0x23}
	fmt.Println(myStruct)
	// When String() is NOT defined => {Yuto 35}
	// When String() is defined     => { name: Yuto, age: 35 }
	fmt.Println()

	numberStr := "0123456789"
	slices := numberStr[2:8]
	fmt.Println(slices) // 234567
}
