package utils

import (
	"fmt"
	"reflect"
)

type mixed struct {
	boolValue  bool
	intValue   int32
	intPointer *int
	person     *person
	arrayValue []int
	mapValue   map[string]int
}

func getPropertyData() {
	person := person{
		Name:   "Yuto",
		Age:    99,
		Gender: "Male",
	}

	fmt.Println("--- no pointer value ---")

	value := reflect.ValueOf(person)
	for i := 0; i < value.Type().NumField(); i++ {
		field := value.Type().Field(i)
		fmt.Printf("Key: %s, \ttype: %s, \toffset: %v\n", field.Name, field.Type, field.Offset)
	}

	fmt.Println("--- pointer value ---")

	pValue := reflect.ValueOf(&person)
	for i := 0; i < pValue.Elem().Type().NumField(); i++ {
		field := pValue.Elem().Type().Field(i)
		fmt.Printf("Key: %s, \ttype: %s, \toffset: %v\n", field.Name, field.Type, field.Offset)
	}
}

func getValues() {
	person := person{
		Name:   "Yuto",
		Age:    99,
		Gender: "Male",
	}

	fmt.Println("--- no pointer value ---")
	fmt.Println(">> by Field()")

	value := reflect.ValueOf(person)
	for i := 0; i < value.Type().NumField(); i++ {
		field := value.Field(i)
		fmt.Printf("type: %s, \tvalue: %v\n", field.Type().Name(), field.Interface())
	}

	fmt.Println(">> by FieldByName()")
	name := value.FieldByName("Name")
	age := value.FieldByName("Age")
	fmt.Printf("Name: %v, Age: %v\n", name.Interface(), age.Interface())

	fmt.Println("--- pointer value ---")

	pValue := reflect.ValueOf(&person)
	for i := 0; i < pValue.Elem().Type().NumField(); i++ {
		field := pValue.Elem().Field(i)
		switch field.Type().Kind() {
		case reflect.String:
			fmt.Printf("string data: %s\n", field.String())
		case reflect.Int:
			fmt.Printf("int data: %d\n", field.Int())
		default:
			fmt.Printf("interface data: %v\n", field.Interface())
		}
	}
}

func getKeyAndValue() {
	person := person{
		Name:   "Yuto",
		Age:    99,
		Gender: "Male",
	}

	value := reflect.ValueOf(person)
	for i := 0; i < value.Type().NumField(); i++ {
		typeField := value.Type().Field(i)
		data := value.Field(i).Interface()

		fmt.Printf("Key: %s, \ttype: %s, \tvalue: %v\n", typeField.Name, typeField.Type, data)
	}
}

func handleComplexStruct() {
	intValue := 99
	data := mixed{
		boolValue:  true,
		intValue:   55,
		intPointer: &intValue,
		person: &person{
			Name:   "Yuto",
			Age:    55,
			Gender: "Male",
		},
		arrayValue: []int{1, 2, 3, 4},
		mapValue:   map[string]int{"prop1": 11, "prop2": 22},
	}

	value := reflect.ValueOf(data)
	recursiveValue(value)
}

func recursiveValue(value reflect.Value) {
	valueType := value.Type()
	for i := 0; i < valueType.NumField(); i++ {
		field := value.Field(i)
		switch field.Kind() {
		case reflect.Bool:
			fmt.Printf("bool value (%s: %v)\n", valueType.Field(i).Name, field.Bool())
		case reflect.Int, reflect.Int32:
			fmt.Printf("int value (%s: %v)\n", valueType.Field(i).Name, field.Int())
		case reflect.String:
			fmt.Printf("string value (%s: %v)\n", valueType.Field(i).Name, field.String())
		case reflect.Pointer:
			if field.Elem().Kind() == reflect.Struct {
				recursiveValue(field.Elem())
			} else {
				fmt.Printf("pointer value. type: %s, (%s: %v)\n", field.Elem().Type().Name(), valueType.Field(i).Name, field.Elem())
			}
		case reflect.Array, reflect.Slice:
			fmt.Printf("array/slice value (%s: %v)\n", valueType.Field(i).Name, field.Slice(0, field.Len()))
		case reflect.Map:
			fmt.Printf("map value. Key: %s\n", valueType.Field(i).Name)
			iter := field.MapRange()
			for iter.Next() {
				fmt.Printf("  %s: %v\n", iter.Key(), iter.Value())
			}
		default:
			fmt.Printf("unsupported type: %s\n", field.Kind().String())
		}
	}
}

func RunDynamicAccess() {
	fmt.Println("=== property data ===")
	getPropertyData()

	fmt.Println("=== values ===")
	getValues()

	fmt.Println("=== key and value ===")
	getKeyAndValue()

	fmt.Println("=== complicated struct ===")
	handleComplexStruct()
}
