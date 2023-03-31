package state

import (
	"errors"
)

type State string

const (
	Running State = "Running"
	Stop    State = "Stop"
	Ready   State = "Ready"
	Error   State = "Error"
)

func IsState(value State) bool {
	states := []State{Running, Stop, Ready, Error}
	for _, state := range states {
		if state == value {
			return true
		}
	}
	return false
}

type MyState struct {
	key   string
	value int
}

func (state MyState) String() string {
	return state.key
}

func (state MyState) GetValue() string {
	return state.key
}
func (state *MyState) GetID() int {
	return state.value
}

func (MyState) Running() MyState {
	return MyState{key: "Running", value: 0}
}
func (MyState) Stop() MyState {
	return MyState{key: "Stop", value: 1}
}
func (MyState) Ready() MyState {
	return MyState{key: "Ready", value: 2}
}
func (MyState) Error() MyState {
	return MyState{key: "Error", value: 3}
}

var errNotMyState = errors.New("key is not MyState")

func (MyState) GetValueByKey(key string) (int, error) {
	switch key {
	case "Running":
		return MyState{}.Running().value, nil
	case "Stop":
		return MyState{}.Stop().value, nil
	case "Ready":
		return MyState{}.Ready().value, nil
	case "Error":
		return MyState{}.Error().value, nil
	default:
		return 0, errNotMyState
	}
}
func (MyState) GetKeyByValue(value int) (string, error) {
	switch value {
	case 0:
		return MyState{}.Running().key, nil
	case 1:
		return MyState{}.Stop().key, nil
	case 2:
		return MyState{}.Ready().key, nil
	case 3:
		return MyState{}.Error().key, nil
	default:
		return "", errNotMyState
	}
}

type Weather int

const (
	Cloudy Weather = iota
	Sunny
	Rainy
	Snowy
)

type Product int

const (
	Notebook Product = iota
	Mouse
	UsbStick
)

func IsWeather(value Weather) bool {
	weathers := []Weather{Cloudy, Sunny, Rainy, Snowy}
	for _, weather := range weathers {
		if weather == value {
			return true
		}
	}
	return false
}

func (w Weather) String() string {
	switch w {
	case Cloudy:
		return "Cloudy"
	case Sunny:
		return "Sunny"
	case Rainy:
		return "Rainy"
	case Snowy:
		return "Snowy"
	default:
		return ""
	}
}
