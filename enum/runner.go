package enum

import (
	"fmt"
	"play-with-go-lang/enum/state"
)

func requireEnum(value state.State) {
	if !state.IsState(value) {
		fmt.Println("not State value")
		return
	}
	fmt.Println(value)

}

func RunEnum() {
	requireEnum(state.Error)
	requireEnum(state.Ready)
	requireEnum("Ready")
	requireEnum("non-enum")
}
