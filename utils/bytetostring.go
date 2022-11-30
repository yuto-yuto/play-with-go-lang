package utils

import (
	"bytes"
	"fmt"
)

func ConvertByteToString() {
	rawBytes := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58}
	fmt.Printf("%s\n", rawBytes) // 0123456789

	fmt.Printf("%s\n", string(rawBytes)) // 0123456789
	fmt.Printf("%s\n", bytes.NewBuffer(rawBytes).String()) // 0123456789
	
	rawByteString := fmt.Sprint(rawBytes)
	fmt.Printf("%s\n", rawByteString)	// [48 49 50 51 52 53 54 55 56 57 58]
}
