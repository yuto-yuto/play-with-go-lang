package utils

import "fmt"

func SliceTest(){
	var list []int

	fmt.Println(list)
	var emptyList []int
	list = append(list, emptyList...)
	fmt.Println(list)
}