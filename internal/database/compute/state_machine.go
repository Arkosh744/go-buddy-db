package compute

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Arkosh744/go-buddy-db/pkg/models"
)

const (
	initialState = iota
	charFoundState
	whiteSpaceFoundState
)

var errNotValidTransition = errors.New("not valid transition")

type transitionAction func(ch rune, stringBuffer *strings.Builder, tokens []string) ([]string, int)

// transition[fromState][ToState].
var transitions = map[int]map[int]transitionAction{
	initialState: {
		charFoundState:       handleCharFound,
		whiteSpaceFoundState: handleWhiteSpaceFound,
	},
	charFoundState: {
		charFoundState:       handleCharFound,
		whiteSpaceFoundState: handleWhiteSpaceFound,
	},
	whiteSpaceFoundState: {
		charFoundState:       handleCharFound,
		whiteSpaceFoundState: handleWhiteSpaceFound,
	},
}

type stateMachine struct {
	state int

	transitions  map[int]map[int]transitionAction
	stringBuffer strings.Builder
	tokens       []string
}

func newStateMachine() *stateMachine {
	return &stateMachine{
		state:        initialState,
		transitions:  transitions,
		stringBuffer: strings.Builder{},
		tokens:       []string{},
	}
}

func (sm *stateMachine) parse(query string) ([]string, error) {
	for _, ch := range query {
		toState, err := getCharType(ch)
		if err != nil {
			return nil, err
		}

		transition, ok := sm.transitions[sm.state][toState]
		if !ok {
			return nil, errNotValidTransition
		}

		sm.tokens, sm.state = transition(ch, &sm.stringBuffer, sm.tokens)
	}

	if action, exists := sm.transitions[sm.state][whiteSpaceFoundState]; exists {
		sm.tokens, _ = action(' ', &sm.stringBuffer, sm.tokens)
	}

	return sm.tokens, nil
}

func getCharType(ch rune) (int, error) {
	switch {
	case isValidCharacter(ch):
		return charFoundState, nil
	case isSpace(ch):
		return whiteSpaceFoundState, nil
	}

	return 0, models.ErrUnknownCharacter
}

func handleCharFound(ch rune, stringBuffer *strings.Builder, tokens []string) ([]string, int) {
	stringBuffer.WriteRune(ch)
	return tokens, charFoundState
}

func handleWhiteSpaceFound(_ rune, stringBuffer *strings.Builder, tokens []string) ([]string, int) {
	if stringBuffer.Len() > 0 {
		tokens = append(tokens, stringBuffer.String())
		stringBuffer.Reset()
	}

	return tokens, whiteSpaceFoundState
}

func isSpace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func isValidCharacter(ch rune) bool {
	return unicode.IsLetter(ch) ||
		unicode.IsDigit(ch) ||
		ch == '_' ||
		ch == '/' ||
		ch == '*'
}
