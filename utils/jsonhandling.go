package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type department struct {
	Name    string     `json:"name"`
	Members []employee `json:"members,omitempty`
}

type employee struct {
	Name          string "json:name"
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	Job           string `json:"role"`
	WithoutSchema string
}

type person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"fooo"`
}

type jsonTestString struct {
	StringPointer0 *string `json:"stringPointer0"`
	StringPointer1 *string `json:"stringPointer1"`
	StringPointer2 *string `json:"stringPointer2,omitempty"`
	StringValue    string  `json:"stringValue,omitempty"`
}

type jsonTestDefault struct {
	IntValue  int            `json:"intValue,omitempty"`
	IntArray  []int          `json:"intArray,omitempty"`
	BoolValue bool           `json:"boolValue,omitempty"`
	MapValue  map[int]string `json:"mapValue,omitempty"`
}

type jsonTestOther struct {
	AsIs    int    `json:""`
	Ignored string `json:"-"`
	Dash    string `json:"-,"`
}

const (
	prefix = ""
	indent = "\t"
)

func convertBetweenStructAndJson() {
	type myProduct struct {
		Name  string
		Price int
	}

	item := myProduct{
		Name:  "Apple",
		Price: 55,
	}
	jsonBytes, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
	// {"Name":"Apple","Price":55}

	jsonBytes2, _ := json.MarshalIndent(item, prefix, indent)
	fmt.Println(string(jsonBytes2))

	var fromJson []myProduct
	myJson := []byte(`[
		{"Name": "Apple", "prICe": 11, "id": 3},
		{"Name": "Melon", "PricE": 22, "id": 4}
	]`)
	err = json.Unmarshal(myJson, &fromJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", fromJson)
	// [{Name:Apple Price:11} {Name:Melon Price:22}]
}

func basic() {
	employee := employee{
		Name:   "Yuto",
		Age:    35,
		Gender: "Male",
		Job:    "Software Developer",
	}
	jsonBytes, _ := json.Marshal(employee)
	// fmt.Println(jsonBytes)
	// fmt.Println(string(jsonBytes))

	jsonBytes2, _ := json.MarshalIndent(employee, prefix, indent)
	fmt.Println(string(jsonBytes2))

	fmt.Println("----- Unmarshal -----")
	var yuto person
	err := json.Unmarshal(jsonBytes, &yuto)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", yuto)
	fmt.Printf("Name: %s, Age: %d, Gender: %s\n", yuto.Name, yuto.Age, yuto.Gender)
}

func nullAssign() {
	str := ""
	obj := jsonTestString{
		StringPointer0: &str,
		StringPointer1: nil,
	}
	jsonBytes, _ := json.MarshalIndent(obj, prefix, indent)
	fmt.Println(string(jsonBytes))
	// {
	//     "stringPointer0": "",
	//     "stringPointer1": null,
	// }
}

func ignoreKey() {
	obj := jsonTestOther{
		AsIs:    11,
		Ignored: "this is ignored",
		Dash:    "this is not ignored",
	}
	jsonBytes, _ := json.MarshalIndent(obj, prefix, indent)
	fmt.Println(string(jsonBytes))
	// {
	//     "AsIs": 11,
	//     "-": "this is not ignored"
	// }
}

func runWithEmpty() {
	{
		str := ""
		obj := jsonTestString{
			StringPointer2: &str,
			StringValue:    "",
		}
		jsonBytes, _ := json.MarshalIndent(obj, prefix, indent)
		fmt.Println(string(jsonBytes))
		// {
		//     "stringPointer0": null,
		//     "stringPointer1": null,
		//     "stringPointer2": "",
		// }
	}
	{
		obj := jsonTestDefault{
			IntValue:  0,
			IntArray:  []int{},
			BoolValue: false,
			MapValue:  map[int]string{},
		}
		jsonBytes, _ := json.MarshalIndent(obj, prefix, indent)
		fmt.Println(string(jsonBytes))
		// {}
	}
	{
		obj := jsonTestDefault{
			IntValue:  -1,
			IntArray:  []int{0},
			BoolValue: true,
			MapValue:  map[int]string{0: ""},
		}
		jsonBytes, _ := json.MarshalIndent(obj, prefix, indent)
		fmt.Println(string(jsonBytes))
		// {
		//         "intValue": -1,
		//         "intArray": [
		//                 0
		//         ],
		//         "boolValue": true,
		//         "mapValue": {
		//                 "0": ""
		//         }
		// }
	}
}

func nested() {
	employees := []employee{
		{
			Name:   "Yuto",
			Age:    35,
			Gender: "Male",
			Job:    "Software Developer",
		},
		{
			Name:   "Daniel",
			Age:    40,
			Gender: "Male",
			Job:    "Software Tester",
		},
	}

	department := department{
		Name:    "Technical Feeder",
		Members: employees,
	}
	jsonBytes, _ := json.MarshalIndent(department, prefix, indent)
	fmt.Println(string(jsonBytes))
	// {
	//     "name": "Technical Feeder",
	//     "Members": [
	//             {
	//                     "Name": "Yuto",
	//                     "age": 35,
	//                     "gender": "Male",
	//                     "role": "Software Developer",
	//                     "WithoutSchema": ""
	//             },
	//             {
	//                     "Name": "Daniel",
	//                     "age": 40,
	//                     "gender": "Male",
	//                     "role": "Software Tester",
	//                     "WithoutSchema": ""
	//             }
	//     ]
	// }
}

func showData(data any) {
	dataType := reflect.TypeOf(data).Kind().String()
	fmt.Printf("type: %s, value: %+v\n", dataType, data)
}

func withoutStruct() {
	var objMap map[string]any

	if err := json.Unmarshal([]byte(`{
		"name": "Yuto",
		"prop1": 12,
		"prop2": 2.2,
		"prop3": false,
		"prop4": null,
		"prop5": [1,2,3,"str"],
		"prop6": {
			"prop6-nested": 15
		}
		}`), &objMap); err != nil {
		fmt.Println(err)
		return
	}
	showData(objMap)                            // type: map, value: map[name:Yuto prop1:12 prop2:2.2 prop3:false prop4:<nil> prop5:[1 2 3 str] prop6:map[prop6-nested:15]]
	showData(objMap["name"])                    // type: string, value: Yuto
	showData(objMap["prop1"])                   // type: float64, value: 12
	fmt.Println(objMap["prop1"] == 12)          // false
	fmt.Println(objMap["prop1"] == float64(12)) // true
	showData(objMap["prop2"])                   // type: float64, value: 2.2
	showData(objMap["prop3"])                   // type: bool, value: false
	fmt.Println(objMap["prop4"])                // <nil>

	array := objMap["prop5"]
	showData(array) // type: slice, value: [1 2 3,str]
	handleArrayAndMap(array)
	// type: float64, value: 1
	// type: float64, value: 2
	// type: float64, value: 3
	// type: string, value: str

	valueOfArray := reflect.ValueOf(array)
	showData(valueOfArray.Index(0).Interface())     // type: float64, value: 1
	showData(valueOfArray.Index(0).Elem().Float())  // type: float64, value: 1
	showData(valueOfArray.Index(1).Interface())     // type: float64, value: 2
	showData(valueOfArray.Index(1).Elem().Float())  // type: float64, value: 2
	showData(valueOfArray.Index(2).Interface())     // type: float64, value: 3
	showData(valueOfArray.Index(2).Elem().Float())  // type: float64, value: 3
	showData(valueOfArray.Index(3).Interface())     // type: string, value: str
	showData(valueOfArray.Index(3).Elem().String()) // type: string, value: str

	fmt.Println()

	nestedObj := objMap["prop6"]
	showData(nestedObj)          // type: map, value: map[prop6-nested:15]
	handleArrayAndMap(nestedObj) // key: prop6-nested, value: 15

	valueOfNestedObj := reflect.ValueOf(nestedObj)
	showData(valueOfNestedObj)                                           // type: struct, value: map[prop6-nested:15]
	showData(valueOfNestedObj.MapIndex(reflect.ValueOf("prop6-nested"))) // type: float64, value: 15

	convertedMap, ok := nestedObj.(map[string]any)
	if !ok {
		fmt.Println("conversion error for map")
	} else {
		showData(convertedMap["prop6-nested"]) // type: float64, value: 15
	}

	fmt.Println()
	for key, value := range objMap {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
	// key: prop1, value: %!s(float64=12)
	// key: prop2, value: %!s(float64=2.2)
	// key: prop3, value: %!s(bool=false)
	// key: prop4, value: %!s(<nil>)
	// key: prop5, value: [%!s(float64=1) %!s(float64=2) %!s(float64=3) str]
	// key: prop6, value: map[prop6-nested:%!s(float64=15)]
	// key: name, value: Yuto

	var objMapList []map[string]any

	if err := json.Unmarshal([]byte(`[{"name": "Yuto", "prop1": 12}, {"name": "Yuto2", "prop1": 33}]`), &objMapList); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(objMapList[0]) // map[name:Yuto prop1:12]
	fmt.Println(objMapList[1]) // map[name:Yuto2 prop1:33]
}

func handleArrayAndMap(array any) {
	fmt.Println("--- handleArrayAndMap start ---")
	switch v := array.(type) {
	case []any:
		for _, data := range v {
			showData(data)
		}
	case map[string]any:
		for key, value := range v {
			fmt.Printf("key: %s, value: %+v\n", key, value)
		}
	default:
		fmt.Println("not an array/map")
	}
	fmt.Println("--- handleArrayAndMap end ---")
}

func HandleJson() {
	// convertBetweenStructAndJson()
	// basic()
	// nullAssign()
	// runWithEmpty()
	// ignoreKey()
	// nested()
	withoutStruct()
}
