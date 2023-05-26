package utils

import (
	"fmt"
	"strconv"
)

type UnsignedInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type UnsignedIntWithTilda interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type UintAlias = uint
type MyUnsignedInt uint

func RunTilda() {
	acceptUint(uint(1))
	acceptUint32(uint32(1))
	acceptUint64(uint64(1))
	acceptAny(uint32(22))    // 22
	acceptAny("sample text") // unsupported data type: string

	fmt.Println()
	acceptUintGenerics(uint(5))
	acceptUintGenerics(uint8(6))
	acceptUintGenerics(uint16(7))
	acceptUintGenerics(uint32(8))
	acceptUintGenerics(uint64(9))
	// acceptUintGenerics("sample")

	fmt.Println()

	acceptUnsignedInt(uint(1))
	acceptUnsignedInt(uint8(1))
	acceptUnsignedInt(uint64(1))
	acceptUnsignedInt(UintAlias(1))
	// MyUnsignedInt does not implement UnsignedInt (possibly missing ~ for uint in constraint UnsignedInt)compilerInvalidTypeArg
	// acceptUnsignedInt(MyUnsignedInt(1))

	fmt.Println()

	acceptUnsignedIntWithTilda(uint(1))
	acceptUnsignedIntWithTilda(uint8(1))
	acceptUnsignedIntWithTilda(uint64(1))
	acceptUnsignedIntWithTilda(MyUnsignedInt(1))

	fmt.Println()

	product := productInTilda{
		Name:   "Desk",
		Height: 15,
		Width:  40.3,
		Length: 100.123,
		Weight: 3.2,
	}
	fmt.Println(product)
	fmt.Println(product.Height)
	fmt.Println(product.Width)
}

func acceptUint(value uint) {
	fmt.Println(value)
}
func acceptUint32(value uint32) {
	fmt.Println(value)
}
func acceptUint64(value uint64) {
	fmt.Println(value)
}

func acceptAny(value any) {
	switch v := value.(type) {
	case uint:
		fmt.Println(v)
	case uint32:
		fmt.Println(v)
	case uint64:
		fmt.Println(v)
	default:
		fmt.Printf("unsupported data type: %T", v)
	}
}

func acceptUintGenerics[T uint | uint8 | uint16 | uint32 | uint64](value T) {
	fmt.Println(value)
}

func acceptUnsignedInt[T UnsignedInt](value T) {
	fmt.Println(value)
}

func acceptUnsignedIntWithTilda[T UnsignedIntWithTilda](value T) {
	fmt.Println(value)
}

type Centimeter float64

func (value Centimeter) String() string {
	return fmt.Sprintf("%s cm", toStringWithTwoDigitsAfterDecimalPoint(value, 1))
}

type Kilogram float64

func (value Kilogram) String() string {
	return fmt.Sprintf("%s kg", toStringWithTwoDigitsAfterDecimalPoint(value, 2))
}

type productInTilda struct {
	Name   string
	Height Centimeter
	Width  Centimeter
	Length Centimeter
	Weight Kilogram
}

func (p productInTilda) String() string {
	return fmt.Sprintf("(Name: %s, Height: %s, Width: %s, Length: %s, Weight: %s)",
		p.Name, p.Height, p.Width, p.Length, p.Weight)
}

func toStringWithTwoDigitsAfterDecimalPoint[T ~float64](value T, digit int) string {
	return strconv.FormatFloat(float64(value), 'f', digit, 64)
}
