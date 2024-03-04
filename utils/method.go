package utils

import "fmt"

type MyEnum int8

const (
	_ MyEnum = iota
	Sleeping
	Running
	Idling
)

func (e MyEnum) String() string {
	switch e {
	case Sleeping:
		return "Sleeping"
	case Running:
		return "Running"
	case Idling:
		return "Idling"
	default:
		return "Unknown"
	}
}

type aliasInt int32

func (m aliasInt) Square() aliasInt {
	return m * m
}

func (m *aliasInt) UpdateToSquare(value int) {
	newValue := aliasInt(value * value)
	*m = newValue
}

type Sensor struct {
	Value    int
	DataList []int
}

func (m *Sensor) UpdateValueWithPointer(value int) {
	m.Value = value
}

func (m Sensor) UpdateValue(value int) {
	m.Value = value
}

func (m *Sensor) addValueWithPointer(value int) {
	m.DataList = append(m.DataList, value)
}

func (m Sensor) addrValue(value int) {
	m.DataList = append(m.DataList, value)
}

func updateDataForPointer(current *Sensor, newValue int) {
	if current != nil {
		current.Value = newValue
	}

}

func updateDataForValue(current Sensor, newValue int) {
	current.Value = newValue
}

func addDataToArrayWithPointer(list *[]int) {
	if list != nil {
		*list = append(*list, 1, 2, 3)
	}
}

func addDataToArray(list []int) {
	list = append(list, 1, 2, 3)
}

func RunMethodTest() {
	dataList := []int{99}
	addDataToArray(dataList)
	fmt.Println(dataList) // [99]
	addDataToArrayWithPointer(&dataList)
	fmt.Println(dataList) // 99 1 2 3

	s := Sensor{Value: 1}
	fmt.Println(s.Value) // 1
	s.UpdateValue(2)
	fmt.Println(s.Value) // 1
	s.UpdateValueWithPointer(3)
	fmt.Println(s.Value) // 3

	fmt.Println(s.DataList) // []
	s.addrValue(10)
	fmt.Println(s.DataList) // []
	s.addValueWithPointer(20)
	fmt.Println(s.DataList) // [20]

	s2 := Sensor{Value: 99}
	updateDataForValue(s2, 999)
	fmt.Println(s2.Value) // 99
	updateDataForPointer(&s2, 88)
	fmt.Println(s2.Value) // 88

	fmt.Println("--- aliasInt ---")

	value := aliasInt(5)
	fmt.Println(value)          // 5
	fmt.Println(value.Square()) // 25
	value.UpdateToSquare(10)
	fmt.Println(value) // 100

	fmt.Println()
	var myEnum MyEnum
	fmt.Println(myEnum) // Unknown
	myEnum = Sleeping
	fmt.Println(myEnum) // Sleeping
	myEnum = Running
	fmt.Println(myEnum) // Running
}
