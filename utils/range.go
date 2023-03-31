package utils

import "fmt"

type Emitter struct {
	callbacks map[string]func()
}

func (e *Emitter) Subscribe(key string, cb func()) {
	if len(e.callbacks) == 0 {
		e.callbacks = map[string]func(){key: cb}
	} else {
		e.callbacks[key] = cb
	}
}

func (e *Emitter) Fire(key string) {
	action, prs := e.callbacks[key]
	if prs {
		action()
	}
}

func RunRange() {
	loopIntArray()
	loopIntArray2()
	loopStructArray()
}

func loopIntArray() {
	fmt.Println("--- loopIntArray ---")
	emitter := Emitter{}
	one := 1
	two := 2
	three := 3
	intArray := []int{one, two, three}

	fmt.Println("--- range --- ")
	for index, value := range intArray {
		fmt.Printf("%d: %v\n", index, value)
		// 0: 1
		// 1: 2
		// 2: 3
		emitter.Subscribe(fmt.Sprint(index), func() {
			fmt.Printf("index: %d, value: %d\n", index, value)
		})
	}

	fmt.Println("--- fire --- ")
	emitter.Fire("0") // index: 2, value: 3
	emitter.Fire("1") // index: 2, value: 3
	emitter.Fire("2") // index: 2, value: 3
}

func loopIntArray2() {
	fmt.Println("--- loopIntArray ---")
	emitter := Emitter{}
	one := 1
	two := 2
	three := 3
	intArray := []int{one, two, three}

	fmt.Println("--- range --- ")
	for index, value := range intArray {
		fmt.Printf("%d: %v\n", index, value)
		// 0: 1
		// 1: 2
		// 2: 3
		indexCopy := index
		valueCopy := value
		emitter.Subscribe(fmt.Sprint(index), func() {
			fmt.Printf("index: %d, value: %d\n", indexCopy, valueCopy)
		})
	}

	fmt.Println("--- fire --- ")
	emitter.Fire("0") // index: 0, value: 1
	emitter.Fire("1") // index: 1, value: 2
	emitter.Fire("2") // index: 2, value: 3
}

func loopStructArray() {
	fmt.Println("--- loopIntPointerArray ---")
	emitter := Emitter{}
	book1 := BaseProductInfo{Name: "name-1", Price: 11}
	book2 := BaseProductInfo{Name: "name-2", Price: 22}
	book3 := BaseProductInfo{Name: "name-3", Price: 33}

	structArray := []BaseProductInfo{book1, book2, book3}

	fmt.Println("--- range --- ")
	for index, value := range structArray {
		fmt.Printf("%d: %+v\n", index, value)

		indexCopy := index
		valueCopy := value
		emitter.Subscribe(fmt.Sprint(index), func() {
			fmt.Printf("index: %d, value: %+v\n", indexCopy, valueCopy)
		})
	}

	fmt.Println("--- fire --- ")
	emitter.Fire("0")
	emitter.Fire("1")
	emitter.Fire("2")
}
