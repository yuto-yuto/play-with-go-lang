package utils

import (
	"encoding/json"
	"fmt"
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

func HandleJson() {
	// convertBetweenStructAndJson()
	// basic()
	// nullAssign()
	// runWithEmpty()
	// ignoreKey()
	nested()
}
