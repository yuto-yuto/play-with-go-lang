package utils

import (
	"fmt"
	"unsafe"
)

type myState1 struct {
	ready     bool
	resume    bool
	running   bool
	aborted   bool
	completed bool
}

type myState2 struct {
	state int
}

func (m *myState2) isReady() bool {
	return m.state&1 > 0
}
func (m *myState2) isResume() bool {
	return m.state&2 > 0
}
func (m *myState2) isRunning() bool {
	return m.state&4 > 0
}
func (m *myState2) isAborted() bool {
	return m.state&8 > 0
}
func (m *myState2) isCompleted() bool {
	return m.state&16 > 0
}
func (m *myState2) setReady() {
	m.state |= 1
}
func (m *myState2) setResume() {
	m.state |= 2
}
func (m *myState2) setRunning() {
	m.state |= 4
}
func (m *myState2) setAborted() {
	m.state |= 8
}
func (m *myState2) setCompleted() {
	m.state |= 16
}
func (m *myState2) reset() {
	m.state = 0
}

func RunFlag() {
	state1 := myState1{}
	state2 := myState2{}
	fmt.Println(unsafe.Sizeof(state1)) // 5
	fmt.Println(unsafe.Sizeof(state2)) // 8
}
