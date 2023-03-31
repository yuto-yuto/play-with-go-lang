package enum

import (
	"fmt"
	"play-with-go-lang/enum/state"
)

func requireState(value state.State) {
	if !state.IsState(value) {
		fmt.Println("not State value")
		return
	}
	fmt.Println(value)
}

func requireMyState(value state.MyState) {
	fmt.Printf("key: %s, value: %d\n", value.GetValue(), value.GetID())
}

func requireWeather(value state.Weather) {
	if !state.IsWeather(value) {
		fmt.Println("not Weather value")
		return
	}
	fmt.Printf("key: %s, value: %d\n", value, value)
}

func RunEnum() {
	requireState(state.Error) // Error
	requireState(state.Ready) // Ready
	requireState("Ready")     // Ready
	requireState("non-enum")  // not State value

	fmt.Println()
	requireWeather(state.Cloudy) // key: Cloudy, value: 0
	requireWeather(state.Sunny)  // key: Sunny, value: 1
	requireWeather(0)            // key: Cloudy, value: 0
	requireWeather(3)            // key: Snowy, value: 3
	requireWeather(4)            // not Weather value

	fmt.Println()
	requireMyState(state.MyState{}.Running()) // key: Running, value: 0
	requireMyState(state.MyState{}.Ready())   // key: Ready, value: 2
	// cannot use "abcde" (untyped string constant) as state.MyState value in argument to requireMyState
	// requireMyState("abcde")
	fmt.Println(state.MyState{}.GetKeyByValue(0))         // Running <nil>
	fmt.Println(state.MyState{}.GetKeyByValue(2))         // Ready <nil>
	fmt.Println(state.MyState{}.GetKeyByValue(5))         //  key is not MyState
	fmt.Println(state.MyState{}.GetValueByKey("Running")) // 0 <nil>
	fmt.Println(state.MyState{}.GetValueByKey("Ready"))   // 2 <nil>
	fmt.Println(state.MyState{}.GetValueByKey("Unknown")) // 0 key is not MyState

	fmt.Println(state.MyState{}.Running()) // Running
}
