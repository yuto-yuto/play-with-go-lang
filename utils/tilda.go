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

type MyUnsignedInt uint

func RunTilda() {
	acceptUnsignedInt(uint(1))
	acceptUnsignedInt(uint8(1))
	acceptUnsignedInt(uint64(1))
	// MyUnsignedInt does not implement UnsignedInt (possibly missing ~ for uint in constraint UnsignedInt)compilerInvalidTypeArg
	// acceptUnsignedInt(MyUnsignedInt(1))

	fmt.Println()

	acceptUnsignedInt(uint(1))
	acceptUnsignedInt(uint8(1))
	acceptUnsignedInt(uint64(1))
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
}

func acceptUnsignedInt[T UnsignedInt](value T) {
	fmt.Println(value)
}

func acceptUnsignedIntWithTilda[T UnsignedIntWithTilda](value T) {
	fmt.Println(value)
}

type Centimeter float64

func (value Centimeter) String() string {
	return fmt.Sprintf("%s cm", toStringWithTwoDigitsAfterDecimalPoint(value))
}

type Kilogram float64

func (value Kilogram) String() string {
	return fmt.Sprintf("%s kg", toStringWithTwoDigitsAfterDecimalPoint(value))
}

type productInTilda struct {
	Name   string
	Height Centimeter
	Width  Centimeter
	Length Centimeter
	Weight Kilogram
}

func (p productInTilda) String() string {
	return fmt.Sprintf("(Name: %s, Height: %s, Width: %s, Length: %s, Weight: %s)", p.Name, p.Height.String(), p.Width.String(), p.Length.String(), p.Weight.String())
}

type Float64Value interface {
	~float32 | ~float64
}

func toStringWithTwoDigitsAfterDecimalPoint[T Float64Value](value T) string {
	return strconv.FormatFloat(float64(value), 'f', 2, 64)
}
