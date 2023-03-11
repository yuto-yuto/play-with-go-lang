package utils

import (
	"fmt"
	"sort"
)

type mapKey struct {
	key1 string
	key2 int
}

func RunMap() {
	map1 := make(map[int]string)
	map1[1] = "one"
	map1[7] = "seven"
	map1[3] = "three"
	map1[4] = "four"
	map1[6] = "six"

	map2 := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	showValue(map1, 1)       // value: [one]
	showValue(map1, 2)       // key does not exist [2]
	showValue(map2, "first") // value: [1]
	showValue(map2, "four")  // key does not exist [four]

	iterate(map1)
	// key: 6, value: six
	// key: 1, value: one
	// key: 3, value: three
	// key: 4, value: four
	iterate(map1)
	// key: 3, value: three
	// key: 4, value: four
	// key: 6, value: six
	// key: 1, value: one

	sortAndIterate(map1)
	// key: 1, value: one
	// key: 3, value: three
	// key: 4, value: four
	// key: 6, value: six
	// key: 7, value: seven

	fmt.Println("--- struct key ----")

	mapWithStruct := make(map[mapKey]string)
	mapKey1 := &mapKey{key1: "First", key2: 1}
	mapKey1_2 := &mapKey{key1: "First", key2: 1}
	mapKey2 := &mapKey{key1: "Second", key2: 2}
	mapKey3 := &mapKey{key1: "Third", key2: 3}
	mapWithStruct[*mapKey1] = "Value-1"
	mapWithStruct[*mapKey2] = "Value-2"
	mapWithStruct[*mapKey3] = "Value-3"
	fmt.Println(mapKey1 == mapKey1_2)                          // false
	fmt.Println(*mapKey1 == *mapKey1_2)                        // true
	showValue(mapWithStruct, *mapKey1)                         // value: [Value-1]
	showValue(mapWithStruct, mapKey{key1: "First", key2: 1})   // value: [Value-1]
	showValue(mapWithStruct, mapKey{key1: "unknown", key2: 1}) // key does not exist [{key1:unknown key2:1}]

	fmt.Println("----struct key 2 ---")

	mapWithStruct2 := make(map[*mapKey]string)
	mapWithStruct2[mapKey1] = "Value-1"
	mapWithStruct2[mapKey2] = "Value-2"
	mapWithStruct2[mapKey3] = "Value-3"
	showValue(mapWithStruct, *mapKey1)                       // value: [Value-1]
	showValue(mapWithStruct, mapKey{key1: "First", key2: 1}) // value: [Value-1]
	_, prs := mapWithStruct2[mapKey1]
	fmt.Println(prs) // true

	_, prs = mapWithStruct2[&mapKey{key1: "unknown", key2: 1}]
	fmt.Println(prs) // false

	fmt.Println("--- map of map ----")
	mapOfMap := make(map[string]map[string]int)
	nestedMapA := make(map[string]int)
	nestedMapA["aa"] = 11
	nestedMapA["ab"] = 12
	nestedMapA["ac"] = 13
	nestedMapB := make(map[string]int)
	nestedMapB["ba"] = 21
	nestedMapB["bb"] = 22
	nestedMapB["bc"] = 23
	mapOfMap["a"] = nestedMapA
	mapOfMap["b"] = nestedMapB

	fmt.Println(mapOfMap) // map[a:map[aa:11 ab:12 ac:13] b:map[ba:21 bb:22 bc:23]]

	fmt.Println("--- map of map 2 ----")
	mapOfMap2 := make(map[string]map[string]int)
	mapOfMap2["a"] = make(map[string]int)
	mapOfMap2["a"]["aa"] = 11
	mapOfMap2["a"]["ab"] = 12
	mapOfMap2["a"]["ac"] = 13

	mapOfMap2["b"] = make(map[string]int)
	mapOfMap2["b"]["ba"] = 21
	mapOfMap2["b"]["bb"] = 22
	mapOfMap2["b"]["bc"] = 23

	fmt.Println(mapOfMap2) // map[a:map[aa:11 ab:12 ac:13] b:map[ba:21 bb:22 bc:23]]

	fmt.Println("--- map of map 3 ----")
	mapOfMap3 := map[string]map[string]int{
		"a": {
			"aa": 11,
			"ab": 12,
			"ac": 13,
		},
		"b": {
			"ba": 21,
			"bb": 22,
			"bc": 23,
		},
	}

	fmt.Println(mapOfMap3) // map[a:map[aa:11 ab:12 ac:13] b:map[ba:21 bb:22 bc:23]]

}

func showValue[K comparable, V int | string](m map[K]V, key K) {
	v, prs := m[key]
	if !prs {
		fmt.Printf("key does not exist [%+v] \n", key)
	} else {
		fmt.Printf("value: [%v] \n", v)
	}
}

func iterate(m map[int]string) {
	for key, val := range m {
		fmt.Printf("key: %d, value: %s\n", key, val)
	}
	fmt.Println()
}

func sortAndIterate(m map[int]string) {
	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("key: %d, value: %s\n", key, m[key])
	}
}
