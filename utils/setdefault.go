package utils

import (
	"errors"
	"fmt"
)

type planType string
type bedType string

const (
	planTypeA planType = "A"
	planTypeB planType = "B"
	planTypeC planType = "C"
)

const (
	singleBed bedType = "single"
	doubleBed bedType = "double"
	queenBed  bedType = "queen"
	kingBed   bedType = "king"
)

type planInfo struct {
	plan           planType
	bedType        *bedType
	persons        *int
	numberOfBeds   *int
	numberOfNights *int
}

func (m planInfo) print() {
	fmt.Printf("plan: %s, ", m.plan)

	if m.bedType == nil {
		fmt.Printf("bedType: nil")
	} else {
		fmt.Printf("bedType: %s", *m.bedType)
	}

	if m.persons == nil {
		fmt.Printf(", persons: nil")
	} else {
		fmt.Printf(", persons: %d", *m.persons)
	}

	if m.numberOfBeds == nil {
		fmt.Printf(", numberOfBeds: nil")
	} else {
		fmt.Printf(", numberOfBeds: %d", *m.numberOfBeds)
	}

	if m.numberOfNights == nil {
		fmt.Printf(", numberOfNights: nil\n")
	} else {
		fmt.Printf(", numberOfNights: %d\n", *m.numberOfNights)
	}
}

type getDefaultValues func() (int, bedType, int, int)

func getDefaultValuesForPlanA() (int, bedType, int, int) {
	person := 1
	bedType := singleBed
	numberOfBeds := 1
	numberOfNights := 1

	return person, bedType, numberOfBeds, numberOfNights
}

func getDefaultValuesForPlanB() (int, bedType, int, int) {
	person := 2
	bedType := doubleBed
	numberOfBeds := 2
	numberOfNights := 1

	return person, bedType, numberOfBeds, numberOfNights
}

func getDefaultValuesForPlanC() (int, bedType, int, int) {
	person := 1
	bedType := queenBed
	numberOfBeds := 1
	numberOfNights := 3

	return person, bedType, numberOfBeds, numberOfNights
}

func SetDefaultValue() {
	fmt.Println("---useFunction---")
	useFunction(planInfo{plan: planTypeA})
	useFunction(planInfo{plan: planTypeB})
	useFunction(planInfo{plan: planTypeC})

	chosenBedType := kingBed
	numberOfPersons := 2
	useFunction(planInfo{plan: planTypeC, bedType: &chosenBedType, persons: &numberOfPersons})

	fmt.Println("---useMap---")

	useMap(planInfo{plan: planTypeA})
	useMap(planInfo{plan: planTypeB})
	useMap(planInfo{plan: planTypeC})
	useMap(planInfo{plan: planTypeC, bedType: &chosenBedType, persons: &numberOfPersons})

	fmt.Println("---useSwitchCase1---")

	useSwitchCase1(planInfo{plan: planTypeA})
	useSwitchCase1(planInfo{plan: planTypeB})
	useSwitchCase1(planInfo{plan: planTypeC})
	useSwitchCase1(planInfo{plan: planTypeC, bedType: &chosenBedType, persons: &numberOfPersons})

	fmt.Println("---useSwitchCase2---")

	useSwitchCase2(planInfo{plan: planTypeA})
	useSwitchCase2(planInfo{plan: planTypeB})
	useSwitchCase2(planInfo{plan: planTypeC})
	useSwitchCase2(planInfo{plan: planTypeC, bedType: &chosenBedType, persons: &numberOfPersons})

	fmt.Println("---useStrategy---")

	useStrategy(planInfo{plan: planTypeA})
	useStrategy(planInfo{plan: planTypeB})
	useStrategy(planInfo{plan: planTypeC})
	useStrategy(planInfo{plan: planTypeC, bedType: &chosenBedType, persons: &numberOfPersons})

	fmt.Println("---useStruct---")

	useStruct(planInfo{plan: planTypeA})
	useStruct(planInfo{plan: planTypeB})
	useStruct(planInfo{plan: planTypeC})
	useStruct(planInfo{plan: planTypeC, bedType: &chosenBedType, persons: &numberOfPersons})
}

func useFunction(planInfo planInfo) {
	var defaultPersons int
	var defaultBedType bedType
	var defaultNumberOfBeds int
	var defaultNumberOfNights int

	switch planInfo.plan {
	case planTypeA:
		defaultPersons, defaultBedType, defaultNumberOfBeds, defaultNumberOfNights = getDefaultValuesForPlanA()
	case planTypeB:
		defaultPersons, defaultBedType, defaultNumberOfBeds, defaultNumberOfNights = getDefaultValuesForPlanB()
	case planTypeC:
		defaultPersons, defaultBedType, defaultNumberOfBeds, defaultNumberOfNights = getDefaultValuesForPlanC()
	default:
	}

	if planInfo.persons == nil {
		planInfo.persons = &defaultPersons
	}

	if planInfo.bedType == nil {
		planInfo.bedType = &defaultBedType
	}

	if planInfo.numberOfBeds == nil {
		planInfo.numberOfBeds = &defaultNumberOfBeds
	}

	if planInfo.numberOfNights == nil {
		planInfo.numberOfNights = &defaultNumberOfNights
	}

	planInfo.print()
}

var defaultValueMap = map[planType]map[string]any{
	planTypeA: {
		"bedType":        singleBed,
		"persons":        1,
		"numberOfBeds":   1,
		"numberOfNights": 1,
	},
	planTypeB: {
		"bedType":        doubleBed,
		"persons":        2,
		"numberOfBeds":   2,
		"numberOfNights": 1,
	},
	planTypeC: {
		"bedType":        queenBed,
		"persons":        1,
		"numberOfBeds":   1,
		"numberOfNights": 3,
	},
}

func useMap(planInfo planInfo) {
	defaultMap, ok := defaultValueMap[planInfo.plan]
	if !ok {
		fmt.Printf("no default for plan %s", planInfo.plan)

		return
	}

	if planInfo.persons == nil {
		if val, ok := defaultMap["persons"]; ok {
			if intValue, err := convertAnyToInt(val); err == nil {
				planInfo.persons = intValue
			}
		}
	}

	if planInfo.bedType == nil {
		if val, ok := defaultMap["bedType"]; ok {
			if bedTypeValue, err := convertAnyToBedType(val); err == nil {
				planInfo.bedType = bedTypeValue
			}
		}
	}

	if planInfo.numberOfBeds == nil {
		if val, ok := defaultMap["numberOfBeds"]; ok {
			if intValue, err := convertAnyToInt(val); err == nil {
				planInfo.numberOfBeds = intValue
			}
		}
	}

	if planInfo.numberOfNights == nil {
		if val, ok := defaultMap["numberOfNights"]; ok {
			if intValue, err := convertAnyToInt(val); err == nil {
				planInfo.numberOfNights = intValue
			}
		}
	}

	planInfo.print()
}

func convertAnyToInt(val any) (*int, error) {
	ret, ok := val.(int)
	if !ok {
		return nil, errors.New("conversion error")
	}

	return &ret, nil
}

func convertAnyToBedType(val any) (*bedType, error) {
	ret, ok := val.(bedType)
	if !ok {
		return nil, errors.New("conversion error")
	}

	return &ret, nil
}

func useSwitchCase1(planInfo planInfo) {
	if planInfo.persons == nil {
		var defaultValue int
		switch planInfo.plan {
		case planTypeA:
			defaultValue = 1
		case planTypeB:
			defaultValue = 2
		case planTypeC:
			defaultValue = 1
		}
		planInfo.persons = &defaultValue
	}

	if planInfo.bedType == nil {
		var defaultValue bedType
		switch planInfo.plan {
		case planTypeA:
			defaultValue = singleBed
		case planTypeB:
			defaultValue = doubleBed
		case planTypeC:
			defaultValue = queenBed
		}
		planInfo.bedType = &defaultValue
	}

	if planInfo.numberOfBeds == nil {
		var defaultValue int
		switch planInfo.plan {
		case planTypeA:
			defaultValue = 1
		case planTypeB:
			defaultValue = 2
		case planTypeC:
			defaultValue = 1
		}
		planInfo.numberOfBeds = &defaultValue
	}

	if planInfo.numberOfNights == nil {
		var defaultValue int
		switch planInfo.plan {
		case planTypeA:
			defaultValue = 1
		case planTypeB:
			defaultValue = 1
		case planTypeC:
			defaultValue = 3
		}
		planInfo.numberOfNights = &defaultValue
	}

	planInfo.print()
}

func useSwitchCase2(planInfo planInfo) {
	var defaultPersons int
	var defaultBedType bedType
	var defaultNumberOfBeds int
	var defaultNumberOfNights int

	switch planInfo.plan {
	case planTypeA:
		defaultPersons = 1
		defaultBedType = singleBed
		defaultNumberOfBeds = 1
		defaultNumberOfNights = 1
	case planTypeB:
		defaultPersons = 2
		defaultBedType = doubleBed
		defaultNumberOfBeds = 2
		defaultNumberOfNights = 1
	case planTypeC:
		defaultPersons = 1
		defaultBedType = queenBed
		defaultNumberOfBeds = 1
		defaultNumberOfNights = 3
	}

	if planInfo.persons == nil {
		planInfo.persons = &defaultPersons
	}

	if planInfo.bedType == nil {
		planInfo.bedType = &defaultBedType
	}

	if planInfo.numberOfBeds == nil {
		planInfo.numberOfBeds = &defaultNumberOfBeds
	}

	if planInfo.numberOfNights == nil {
		planInfo.numberOfNights = &defaultNumberOfNights
	}

	planInfo.print()
}

type getDefaults interface {
	getPersons() int
	getBedType() bedType
	getNumberOfBeds() int
	getNumberOfNights() int
}

type planADefaults struct{}

func (planADefaults) getPersons() int        { return 1 }
func (planADefaults) getBedType() bedType    { return singleBed }
func (planADefaults) getNumberOfBeds() int   { return 1 }
func (planADefaults) getNumberOfNights() int { return 1 }

type planBDefaults struct{}

func (planBDefaults) getPersons() int        { return 2 }
func (planBDefaults) getBedType() bedType    { return singleBed }
func (planBDefaults) getNumberOfBeds() int   { return 2 }
func (planBDefaults) getNumberOfNights() int { return 1 }

type planCDefaults struct{}

func (planCDefaults) getPersons() int        { return 1 }
func (planCDefaults) getBedType() bedType    { return queenBed }
func (planCDefaults) getNumberOfBeds() int   { return 1 }
func (planCDefaults) getNumberOfNights() int { return 3 }

func useStrategy(planInfo planInfo) {
	var defaultsGetter getDefaults
	switch planInfo.plan {
	case planTypeA:
		defaultsGetter = planADefaults{}
	case planTypeB:
		defaultsGetter = planBDefaults{}
	case planTypeC:
		defaultsGetter = planCDefaults{}
	}

	if planInfo.persons == nil {
		val := defaultsGetter.getPersons()
		planInfo.persons = &val
	}

	if planInfo.bedType == nil {
		val := defaultsGetter.getBedType()
		planInfo.bedType = &val
	}

	if planInfo.numberOfBeds == nil {
		val := defaultsGetter.getNumberOfBeds()
		planInfo.numberOfBeds = &val
	}

	if planInfo.numberOfNights == nil {
		val := defaultsGetter.getNumberOfNights()
		planInfo.numberOfNights = &val
	}

	planInfo.print()
}

type planDefaults struct {
	bedType        bedType
	persons        int
	numberOfBeds   int
	numberOfNights int
}

func useStruct(planInfo planInfo) {
	var defaultsGetter planDefaults
	switch planInfo.plan {
	case planTypeA:
		defaultsGetter = planDefaults{
			bedType:        singleBed,
			persons:        1,
			numberOfBeds:   1,
			numberOfNights: 1,
		}
	case planTypeB:
		defaultsGetter = planDefaults{
			bedType:        doubleBed,
			persons:        2,
			numberOfBeds:   2,
			numberOfNights: 1,
		}
	case planTypeC:
		defaultsGetter = planDefaults{
			bedType:        queenBed,
			persons:        1,
			numberOfBeds:   1,
			numberOfNights: 3,
		}
	}

	if planInfo.persons == nil {
		planInfo.persons = &defaultsGetter.persons
	}

	if planInfo.bedType == nil {
		planInfo.bedType = &defaultsGetter.bedType
	}

	if planInfo.numberOfBeds == nil {
		planInfo.numberOfBeds = &defaultsGetter.numberOfBeds
	}

	if planInfo.numberOfNights == nil {
		planInfo.numberOfNights = &defaultsGetter.numberOfNights
	}

	planInfo.print()
}
