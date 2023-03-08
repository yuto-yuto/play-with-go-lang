package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ReplaceString() {
	text := "***\nFirst line;\nSecond line;\nThird line;\n***"
	fmt.Println(text)
	// ***
	// First line;
	// Second line;
	// Third line;
	// ***

	fmt.Println(strings.Replace(text, "*", "", -1))
	// First line;
	// Second line;
	// Third line;

	fmt.Println(strings.Replace(text, "*;", "", -1))
	// ***
	// First line;
	// Second line;
	// Third line;
	// ***

	replacer := strings.NewReplacer("*", "", ";", "", "\n", "")
	fmt.Println(replacer.Replace(text))
	// First lineSecond lineThird line

	replacer2 := strings.NewReplacer("*", "-", ";", ".", "\n", " ")
	fmt.Println(replacer2.Replace(text))
	// --- First line. Second line. Third line. ---
	
	regex := regexp.MustCompile("[*\n;]")
	fmt.Println(regex.ReplaceAllString(text, ""))
	// First lineSecond lineThird line
}
