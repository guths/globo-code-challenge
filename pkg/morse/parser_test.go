package morse_test

import (
	"testing"

	"github.com/guths/globo-code-challenge/pkg/morse"
)

var morser morse.Morser

func getMorser() morse.Morser {
	if morser == nil {
		morser = morse.NewMorser()
	}

	return morser
}

func TestMorseDecode(t *testing.T) {
	morseCode := ".... . .-.. .-.. ---"

	result := getMorser().Decode(morseCode)

	if result != "HELLO" {
		t.Fail()
	}
}

func TestMorseDecodeWithEmptySpaces(t *testing.T) {
	morseCode := " .... . .-.. .-.. --- "

	result := getMorser().Decode(morseCode)

	if result != "HELLO" {
		t.Fail()
	}
}

func TestMorseDecodeWithNonMorseCode(t *testing.T) {
	morseCode := "HELLO"

	result := getMorser().Decode(morseCode)

	if result != "" {
		t.Fail()
	}
}

func TestMorseDecodeWithEmptyCode(t *testing.T) {
	morseCode := ""

	result := getMorser().Decode(morseCode)

	if result != "" {
		t.Fail()
	}
}
