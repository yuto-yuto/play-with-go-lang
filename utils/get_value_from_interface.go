package utils

import (
	"fmt"
	"reflect"
)

type unknownObj1 struct {
	Name  string
	Prop1 string
}

type unknownObj2 struct {
	Obj   unknownObj1
	Count int
	Name  string
	Prop1 string
}

func getValueOf(obj any, field string) any {
	value := reflect.ValueOf(obj)
	val := reflect.Indirect(value).FieldByName(field)
	if val.IsValid() {
		return val
	}
	return nil
}

func acceptAnyValue(arg any) {
	switch v := arg.(type) {
	case string:
		fmt.Printf("String: %s", v)
	case int:
		fmt.Printf("Int32: %d", v)
	case float64:
		fmt.Printf("float64: %f", v)
	case unknownObj1:
		fmt.Println(reflect.TypeOf(arg) == reflect.TypeOf(unknownObj1{})) // true
		fmt.Println(reflect.TypeOf(arg) == reflect.TypeOf(unknownObj2{})) // false
		fmt.Printf("Name: %s, Prop1: %s", v.Name, v.Prop1)
	case unknownObj2:
		fmt.Println(reflect.TypeOf(arg) == reflect.TypeOf(unknownObj1{})) // false
		fmt.Println(reflect.TypeOf(arg) == reflect.TypeOf(unknownObj2{})) // true
		fmt.Printf("Name: %s, Prop1: %s, Count: %d", v.Obj.Name, v.Obj.Prop1, v.Count)
	default:
		fmt.Printf("Undefined type: %s", reflect.TypeOf(v))
	}

	fmt.Println()
}

func acceptUnknownType(arg any) {
	fmt.Println(getValueOf(arg, "Name"))
	fmt.Println(getValueOf(arg, "Prop1"))
	fmt.Println(getValueOf(arg, "UnknownProp"))
}

func GetValueFromInterface() {
	obj1 := unknownObj1{
		Name: "Yuto", Prop1: "Hello World!",
	}
	obj2 := unknownObj2{
		Obj: unknownObj1{
			Name: "Agent", Prop1: "ABCDE",
		},
		Count: 66,
		Name:  "Yuto-2",
		Prop1: "Level 1",
	}
	acceptAnyValue(obj1)
	acceptAnyValue(obj2)
	acceptAnyValue(523)
	acceptAnyValue(11.123)
	acceptAnyValue("value")
	acceptAnyValue(make(map[string]int))
	acceptUnknownType(obj1)
	acceptUnknownType(obj2)
}
