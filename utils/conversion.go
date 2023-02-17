package utils

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Bit32 = 32
	Bit64 = 64
	DEC   = 10
	HEX   = 16
	BI    = 2
)

func ConvertBetweenStringAndNumber() {
	fmt.Println("--- int -> string ---")
	intValue := 128
	fmt.Println(strconv.FormatInt(int64(intValue), DEC)) // 128
	fmt.Println(strconv.FormatInt(int64(intValue), HEX)) // 80
	fmt.Println(strconv.FormatInt(int64(intValue), BI))  // 10000000
	fmt.Println(strconv.Itoa(intValue))                  // 128
	fmt.Println(fmt.Sprint(intValue))                    // 128

	result := fmt.Sprintf("%d", intValue)
	fmt.Println(result) // 128
	result = fmt.Sprint(intValue)
	fmt.Println(result) // 128

	fmt.Println("--- int -> float64 ---")
	fmt.Println(float64(intValue)) // 128
	fmt.Println(float32(intValue)) // 128
	// fmt.Printf format %f has arg intValue of wrong type int
	fmt.Printf("%f\n", intValue)          // %!f(int=128)
	fmt.Printf("%f\n", float64(intValue)) // 128.000000
	fmt.Printf("%f\n", float32(intValue)) // 128.000000

	fmt.Println("--- int -> int64 ---")
	fmt.Println(int64(intValue))  // 128
	fmt.Println(uint64(intValue)) // 128
	intNegative := -1
	fmt.Println(uint64(intNegative))     // 18446744073709551615
	fmt.Println(uint64(intNegative + 1)) // 0

	fmt.Println("--- string -> int32/int64 ---")
	strValue := "128"
	fmt.Println(strconv.ParseInt(strValue, DEC, Bit32))  // 128 <nil>
	fmt.Println(strconv.ParseInt(strValue, DEC, Bit64))  // 128 <nil>
	fmt.Println(strconv.ParseInt(strValue, HEX, Bit32))  // 296 <nil>
	fmt.Println(strconv.ParseInt(strValue, HEX, Bit64))  // 296 <nil>
	fmt.Println(strconv.ParseInt(strValue, BI, Bit32))   // 0 strconv.ParseInt: parsing "128": invalid syntax
	fmt.Println(strconv.ParseInt(strValue, BI, Bit64))   // 0 strconv.ParseInt: parsing "128": invalid syntax
	fmt.Println(strconv.ParseInt("11110000", BI, Bit32)) // 240 <nil>
	fmt.Println(strconv.ParseInt("11110000", BI, Bit64)) // 240 <nil>

	fmt.Println("--- string -> float64 ---")
	strFloatValue := "128.123"
	fmt.Println(strconv.ParseFloat(strFloatValue, Bit32)) // 128.1230010986328 <nil>
	fmt.Println(strconv.ParseFloat(strFloatValue, Bit64)) // 128.123 <nil>

	fmt.Println("--- string -> int ---")
	result2, _ := strconv.ParseInt(strValue, DEC, Bit64)
	fmt.Println(int(result2))           // 128
	fmt.Println(strconv.Atoi(strValue)) // 128 <nil>

	fmt.Println("--- float64 -> string ---")
	floatValue := 128.123456
	fmt.Println(strconv.FormatFloat(floatValue, 'f', -1, Bit64)) // 128.123456
	fmt.Println(strconv.FormatFloat(floatValue, 'f', 0, Bit64))  // 128
	fmt.Println(strconv.FormatFloat(floatValue, 'f', 2, Bit64))  // 128.12
	fmt.Println(strconv.FormatFloat(floatValue, 'f', 5, Bit64))  // 128.12346
	fmt.Println(strconv.FormatFloat(floatValue, 'f', 10, Bit64)) // 128.1234560000

	fmt.Println(strconv.FormatFloat(floatValue, 'b', -1, Bit64)) // 4507943349211095p-45
	fmt.Println(strconv.FormatFloat(floatValue, 'e', -1, Bit64)) // 1.28123456e+02
	fmt.Println(strconv.FormatFloat(floatValue, 'E', -1, Bit64)) // 1.28123456E+02
	fmt.Println(strconv.FormatFloat(floatValue, 'g', -1, Bit64)) // 128.123456
	fmt.Println(strconv.FormatFloat(floatValue, 'G', -1, Bit64)) // 128.123456
	fmt.Println(strconv.FormatFloat(floatValue, 'x', -1, Bit64)) // 0x1.003f359ff4fd7p+07
	fmt.Println(strconv.FormatFloat(floatValue, 'X', -1, Bit64)) // 0X1.003F359FF4FD7P+07

	fmt.Println("--- float64 -> int ---")
	fmt.Println(int(floatValue))   // 128
	fmt.Println(int8(floatValue))  // -128
	fmt.Println(int16(floatValue)) // 128
	fmt.Println(int32(floatValue)) // 128
	fmt.Println(int64(floatValue)) // 128
	floatValue2 := 128.99999
	fmt.Println(int(floatValue2))         // 128
	fmt.Println(int(math.Round(128.599))) // 129
	fmt.Println(int(math.Round(128.499))) // 128
	fmt.Println(int(math.Ceil(128.001)))  // 129
	fmt.Println(int(math.Trunc(128.999))) // 128
}
