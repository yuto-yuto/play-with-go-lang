package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func showResult(list []string) {
	fmt.Printf("length: %d, %v\n", len(list), list)
}

func SplitString() {
	fmt.Println("--- strings.Split ---")
	showResult(strings.Split("aaa,bbb,cc,dd", ","))
	showResult(strings.Split("aaa,bbb,cc, dd", ", "))

	fmt.Println("--- regexp.MustCompile ---")
	regex := regexp.MustCompile("[,:]")
	showResult(regex.Split("a:aa:aaa,b:bb:bbb,c:cc:ccc,d:dd:ddd", -1))
	showResult(regex.Split("a:aa:aaa,b:bb:bbb,c:cc:ccc,d:dd:ddd", 0))
	showResult(regex.Split("a:aa:aaa,b:bb:bbb,c:cc:ccc,d:dd:ddd", 1))
	showResult(regex.Split("a:aa:aaa,b:bb:bbb,c:cc:ccc,d:dd:ddd", 3))

	fmt.Println("--- String to number ---")
	numberStr := "1.21"
	integerPartStr := strings.Split(numberStr, ".")[0]
	fmt.Printf("type: %T, value: %s\n", integerPartStr, integerPartStr)
	floatPartStr := strings.Split(numberStr, ".")[1]
	fmt.Printf("type: %T, value: %s\n", floatPartStr, floatPartStr)

	integerPartNum, err := strconv.Atoi(integerPartStr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("type: %T, value: %d\n", integerPartNum, integerPartNum)
	}
	floatPartNum, err := strconv.Atoi(floatPartStr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("type: %T, value: %d\n", floatPartNum, floatPartNum)
	}
}
