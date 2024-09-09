package morse

import (
	"strings"
)

var morseToAlpha = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D",
	".": "E", "..-.": "F", "--.": "G", "....": "H",
	"..": "I", ".---": "J", "-.-": "K", ".-..": "L",
	"--": "M", "-.": "N", "---": "O", ".--.": "P",
	"--.-": "Q", ".-.": "R", "...": "S", "-": "T",
	"..-": "U", "...-": "V", ".--": "W", "-..-": "X",
	"-.--": "Y", "--..": "Z",
	"-----": "0", ".----": "1", "..---": "2", "...--": "3",
	"....-": "4", ".....": "5", "-....": "6", "--...": "7",
	"---..": "8", "----.": "9",
	".-.-.-": ".", ",--..": ",", "?..--..": "?", "'": "'",
	"-.-.--": "!", "/--.": "/", "(": "(", "-.--.-": ")",
	"&": "&", ":---:": ":", ";-.-.": ";", "=": "=",
	"+": "+", ".-.-.": ".", "/": "/",
	"@": "@", " ": " ",
}

type Morser interface {
	Decode(word string) string
}

type morser struct{}

func NewMorser() *morser {
	return &morser{}
}

func (m *morser) Decode(word string) string {
	var builder strings.Builder
	morseChars := strings.Split(word, " ")

	for _, mc := range morseChars {
		converted := morseToAlpha[mc]
		builder.WriteString(converted)
	}

	return builder.String()
}
