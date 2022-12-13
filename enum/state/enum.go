package state

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
