package utils

import (
	"fmt"
	"reflect"
)

type BaseProductInfo struct {
	Name  string
	Price int
}

type Book struct {
	BaseProductInfo
	Taste string
}

type BookHolder struct {
	book        Book
	bookPointer *Book
}

func DoStructTest() {
	book := Book{
		BaseProductInfo: BaseProductInfo{
			Name:  "ABC",
			Price: 11,
		},
		Taste: "Hoo",
	}

	bookPointer := Book{
		BaseProductInfo: BaseProductInfo{
			Name:  "BBB",
			Price: 113,
		},
		Taste: "Foo",
	}

	bookHolder := BookHolder{book: book, bookPointer: &bookPointer}
	fmt.Println(bookHolder.book == book)
	fmt.Println(*bookHolder.bookPointer == bookPointer)
	fmt.Println(reflect.DeepEqual(bookHolder.book, book))
	fmt.Println(reflect.DeepEqual(*bookHolder.bookPointer, bookPointer))
}
