package fsm

const NoNextState = -1

type FSM struct {
	States []int
	InitialState int
	AcceptingStates []int
	NextState func(currentState int, input rune) int
}

func NewFSM(states []int, initialState int, acceptingStates []int, NextState func(currentState int, input rune) int) FSM {
	return FSM{
		states,
		initialState,
		acceptingStates,
		NextState,
	}
}

func (fsm *FSM) Run(input string) (bool, string) {
	currentState := fsm.InitialState
	matched := ""

	for i := 0; i < len(input); i++ {
		char := runeAt(input, i)
		nextState := fsm.NextState(currentState, char);
		// if intInSlice(nextState, fsm.AcceptingStates) {
		//	matched += string(char)
		//	return true, matched
		//}

		if nextState == NoNextState {
			break
		}

		matched += string(char)

		currentState = nextState
	}

	return intInSlice(currentState, fsm.AcceptingStates), matched 
}

func runeAt(str string, idx int) rune {
	runes := []rune(str)
	return runes[idx]
}

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


