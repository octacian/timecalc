package time

import (
	"fmt"
	"os"
)

var DEBUG bool

func init() {
	if os.Getenv("DEBUG") == "TRUE" {
		fmt.Println("Debug mode enabled")
		DEBUG = true
	}
}

// PrintDebug prints a string only if debug mode is enabled
func PrintDebug(str string) {
	if DEBUG {
		fmt.Print(str)
	}
}

// FromString takes a string of user input and returns a response.
func FromString(str string) (string, error) {
	tokens, err := Tokenize(str)
	if DEBUG {
		fmt.Println(Reconstruct(tokens, true))
	}
	if err != nil {
		return "", err
	}

	instructions, err := Parse(tokens)

	if DEBUG {
		fmt.Println(InstructionsToString(instructions))
	}

	if err != nil {
		return "", err
	}

	return Compile(instructions).String(), nil
}
