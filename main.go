package main

import (
	"play-with-go-lang/scanner"
	"play-with-go-lang/utils"
)

func unused(x ...interface{}) {}

func main() {
	unused(scanner.ErrScannerNotFound)

	// scanner.Run()
	// utils.ConvertByteToString()
	utils.StringHandling()
}
