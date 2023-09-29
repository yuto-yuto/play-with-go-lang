package utils

import (
	"encoding/json"
	"fmt"
)

func ToFromJSON() {
	jsonStr := `{
		"num": 123,
		"str": "string-value",
		"boolean": false,
		"nil": null,
		"array": [1, 2, 3],
		"jsonObj": {
			"prop1": 1,
			"prop2": "str-value"
		}
	}`

	fmt.Println(jsonStr)
	var mapValue1 map[string]any

	err := json.Unmarshal([]byte(jsonStr), &mapValue1)
	if err != nil {
		fmt.Printf("failed to unmarshal: %s\n", err.Error())
		return
	} else {
		fmt.Println(mapValue1)
	}

	jsonBytes, err := json.Marshal(mapValue1)
	if err != nil {
		fmt.Printf("failed to marshal: %s\n", err.Error())
		return
	}
	fmt.Println(string(jsonBytes))

	json.Unmarshal(jsonBytes, &mapValue1)
	fmt.Println()
	// mapValue := map[string]any{
	// 	""
	// }
}
