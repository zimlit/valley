package fsm

import (
	"testing"
	"unicode"
)

func TestFSM(t *testing.T) {
	var fsm FSM
	fsm.States = []int{1, 2}
	fsm.InitialState = 1
	fsm.AcceptingStates = []int{2}
	fsm.NextState = func(currentState int, input rune) int {
		switch (currentState) {
		case 1:
			if unicode.IsLetter(input) || input == '_' {
				return 2
			}
		case 2:
			if unicode.IsLetter(input) || unicode.IsDigit(input) || input == '_' {
				return 2;
			}
		}
		return NoNextState
	}

	x, y := fsm.Run("id_3Nt-")

	if x != true || y != "id_3Nt" {
		t.Errorf("%t %s", x, y)
	} 
	t.Log(y)
	
}
