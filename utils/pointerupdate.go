package utils

import (
	"fmt"
	"math"
)

func RunPointerUpdate() {
	value := 0
	updateValue(value)
	fmt.Printf("updateValue: %d\n", value) // 0

	updatePointerValue1(&value)
	fmt.Printf("updatePointerValue1: %d\n", value) // 0
	updatePointerValue2(&value)
	fmt.Printf("updatePointerValue2: %d\n", value) // 99
	updatePointerValue2_2(&value)
	fmt.Printf("updatePointerValue2_2: %d\n", value) // 99

	var pointerValue *int

	updatePointerValue3(&pointerValue)
	fmt.Printf("updatePointerValue3: %d\n", *pointerValue) // 222

	type valueHolder struct {
		value1 *int
		value2 *int
	}

	type floatValueHolder struct {
		value *float64
	}

	var holder1 valueHolder

	holder2 := floatValueHolder{}

	value3 := 3.1
	holder3 := floatValueHolder{
		value: &value3,
	}

	tryToSetValue(&(holder1.value1), holder2.value)
	fmt.Printf("value1: %v, value2: %v\n", holder1.value1, holder1.value2) // value1: <nil>, value2: <nil>
	tryToSetValue(&(holder1.value2), holder3.value)
	fmt.Printf("value1: %v, value2: %v\n", holder1.value1, *holder1.value2) // value1: <nil>, value2: 3
}

func updateValue(target int) {
	target = 11
}

func updatePointerValue1(target *int) {
	a := 99
	target = &a
}

func updatePointerValue2(target *int) {
	*target = 99
}

func updatePointerValue2_2(target *int) {
	value := 100
	target = &value
}

func updatePointerValue3(target **int) {
	a := 222
	*target = &a
}

func tryToSetValue(target **int, from *float64) {
	if from == nil {
		fmt.Println("LOGGING: the data is nil")

		return
	}

	intValue := int(math.Floor(*from))
	*target = &intValue
}
