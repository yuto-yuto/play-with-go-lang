package utils

import (
	"fmt"
	"regexp"
)

func RunRegex() {
	runMatch()
	panicCompile()
	useQuoteMeta()
	fmt.Println()
	runFindSomething()
}

func runMatch() {
	reg, err := regexp.Compile("ID_\\d")
	if err != nil {
		fmt.Println("error in regex string")
	}

	fmt.Println(reg.Match([]byte("ID_1")))   // true
	fmt.Println(reg.Match([]byte("ID_ONE"))) // false

	fmt.Println(reg.MatchString("ID_1"))             // true
	fmt.Println(reg.MatchString("ID_123"))           // true
	fmt.Println(reg.MatchString("something_ID_123")) // true
	fmt.Println(reg.MatchString("ID_ONE"))           // false
}

func runFindSomething() {
	reg, _ := regexp.Compile(`ID_\d`)

	text := "ID_1, ID_42, RAW_ID_52, "
	fmt.Printf("%q\n", reg.FindAllString(text, 2))      // ["ID_1" "ID_4"]
	fmt.Printf("%q\n", reg.FindAllString(text, 5))      // ["ID_1" "ID_4" "ID_5"]
	fmt.Printf("%v\n", reg.FindAllStringIndex(text, 5)) // [[0 4] [6 10] [17 21]]

	result := reg.FindAllString(text, -1)
	fmt.Printf("%q\n", result)               // ["ID_1" "ID_4" "ID_5"]
	fmt.Printf("%q\n", result[2])            // "ID_5"
	fmt.Printf("%q\n", reg.FindString(text)) // "ID_1"

	fmt.Printf("%q\n", reg.FindAllStringSubmatch(text, -1)) // [["ID_1"] ["ID_4"] ["ID_5"]]

	fmt.Println("--- submatch ---")

	reg2, err := regexp.Compile(`([^,\s]+)_(\d+)`)
	if err != nil {
		fmt.Println("error in regex string")
	}

	fmt.Printf("%q\n", reg2.FindAllString(text, -1)) // ["ID_1" "ID_42" "RAW_ID_52"]
	result2 := reg2.FindAllStringSubmatch(text, -1)
	fmt.Printf("%q\n", result2) // [["ID_1" "ID" "1"] ["ID_42" "ID" "42"] ["RAW_ID_52" "RAW_ID" "52"]]
	fmt.Printf("Key: %s, ID: %s\n", result2[2][1], result2[2][2])

	reg3 := regexp.MustCompile(`(\d*)\\*([^\\]+)$`)
	fmt.Printf("%q\n", reg3.FindStringSubmatch("1\\KeyName1"))             // ["1\\KeyName1" "1" "KeyName1"]
	fmt.Printf("%q\n", reg3.FindStringSubmatch("12\\KeyName12"))           // ["12\\KeyName12" "12" "KeyName12"]
	fmt.Printf("%q\n", reg3.FindStringSubmatch("KeyName2"))                // ["KeyName2" "" "KeyName2"]
	fmt.Printf("%q\n", reg3.FindStringSubmatch("something\\12\\KeyName2")) // ["KeyName2" "" "KeyName2"]

	reg4 := regexp.MustCompile(`^([^\\]+)$`)
	fmt.Printf("%q\n", reg4.FindStringSubmatch("KeyName1")) // ["KeyName1" "KeyName1"]

}

func panicCompile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic: ", r)
		}
	}()

	regexp.MustCompile("ID_\\p")
}

func useQuoteMeta() {
	originalStr := "ID_\\p"
	quotedStr := regexp.QuoteMeta(originalStr)
	fmt.Println(originalStr) // ID_\p
	fmt.Println(quotedStr)   // ID_\\p
	_, err := regexp.Compile(quotedStr)
	if err != nil {
		fmt.Println("error in regex string")
	}
}

func runAdvanced() {
	texts := []string{
		`\PRESENT\1\KEY1: 11`,
		`\PRESENT\1\KEY2: "value-2"`,
		`\PRESENT\1\KEY22: ""`,
		`\PRESENT\1\0\KEY3: 22`,
		`\PRESENT\1\0\KEY4: "value-4`,
	}

	// regex := regexp.MustCompile(`(\\PRESENT\\1\\([A-Za-z0-9]+)|\\PRESENT\\1\\[0-9]+\\([A-Za-z0-9]+)): (([0-9]*)|"(.*)")`)
	// regex := regexp.MustCompile(`(\\PRESENT\\1\\([A-Za-z0-9]+)|\\PRESENT\\1\\[0-9]+\\([A-Za-z0-9]+)): ("([^"]*)"|([0-9]*))`)
	// regex := regexp.MustCompile(`\\PRESENT\\1\\([A-Za-z0-9]+): (([0-9]*)|"([^"]*)")`)
	// regex := regexp.MustCompile(`\\PRESENT\\1\\([A-Za-z0-9]+): ("([^"]*)"|([0-9]*))`)
	prefix := `\\PRESENT\\1`
	regexString := fmt.Sprintf(`%s\\[0-9\\]*([A-Za-z0-9]+): "?([^"]*)"?`, prefix)
	regex, err := regexp.Compile(regexString)
	// regex, err := regexp.Compile(`\\PRESENT\\1\\[0-9\\]*([A-Za-z0-9]+): "?([^"]*)"?`)
	if err != nil {
		fmt.Println(regexString)
		fmt.Println(err)
		return
	}

	for _, text := range texts {
		fmt.Println(regex.FindAllString(text, -1))
	}

	for _, text := range texts {
		result := regex.FindStringSubmatch(text)
		fmt.Printf("FindStringSubmatch (len: %d): %+v\n", len(result), result)
		for index, match := range result {
			fmt.Printf("%d: %+v\n", index, match)
		}

		allResult := regex.FindAllStringSubmatch(text, -1)
		fmt.Printf("FindAllStringSubmatch(len: %d): %+v\n", len(allResult), allResult)
		for index, match := range allResult {
			fmt.Printf("%d: %+v\n", index, match)
		}
		fmt.Println()
	}

	pre := `\\\\PRESENT\\\\1`
	// regStr := fmt.Sprintf(`%s\\\\([A-Za-z0-9]+)": "?([^"]*)"?`, pre)
	regStr := fmt.Sprintf(`%s\\\\([0-9]*)\\*([^\\]+)": "?([^"]*)"?`, pre)
	reg := regexp.MustCompile(regStr)
	// fmt.Println(regStr)
	keys := []string{
		"{\n\"\\\\FOLDER\\\\PRESENT\\\\1\\\\NUMBER\": 1",
		"{\n\"\\\\FOLDER\\\\PRESENT\\\\1\\\\0\\\\NUMBER\": 1",
	}
	for _, key := range keys {
		fmt.Println(reg.MatchString(key))
		result := reg.FindStringSubmatch(key)
		fmt.Println(len(result))
		for index, data := range result {
			fmt.Printf("%d: %s\n", index, data)
		}
	}
}
