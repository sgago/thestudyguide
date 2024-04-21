package main

import "fmt"

type State int

type Dfa struct {
	States          map[State]bool
	TransitionTable map[State]map[rune]State
	StartState      State
	AcceptStates    map[State]bool
}

func NewDFA() *Dfa {
	dfa := &Dfa{
		States:          map[State]bool{0: true, 1: true},
		TransitionTable: map[State]map[rune]State{},
		StartState:      0,
		AcceptStates:    map[State]bool{1: true},
	}

	dfa.TransitionTable[0] = map[rune]State{'a': 1}

	return dfa
}

func (dfa *Dfa) Simulate(input string) bool {
	currentState := dfa.StartState

	for _, symbol := range input {
		nextState, ok := dfa.TransitionTable[currentState][symbol]

		if !ok {
			return false // No transition defined for the current symbol
		}

		currentState = nextState
	}

	return dfa.AcceptStates[currentState]
}

func main() {
	dfa := NewDFA()

	inputString := "aaa"

	fmt.Printf("Input \"%s\" is accepted by the DFA: %t\n", inputString, dfa.Simulate(inputString))
}
