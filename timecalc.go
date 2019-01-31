package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// DEBUG defines whether extra verbose information is printed
var DEBUG bool

// PrintDebug prints a string only if debug mode is enabled
func PrintDebug(str string) {
	if DEBUG {
		fmt.Print(str)
	}
}

var OperatorRegex = regexp.MustCompile(`^(\+|\-|\*|\/)$`)
var TimeRegex = regexp.MustCompile(`^([\d][\d]?)?:([\d][\d])?:?(?::([\d][\d]))?$`)
var NumberRegex = regexp.MustCompile(`^((\d+)?\.?)\d+$`)
var WhitespaceRegex = regexp.MustCompile(`^([\s])$`)

func main() {
	if os.Getenv("DEBUG") == "TRUE" {
		DEBUG = true
		fmt.Println("Debug mode enabled.")
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		tokens, err := Tokenize(text)
		if err != nil {
			fmt.Print(err, "\n")
		} else {
			if DEBUG {
				fmt.Println(Reconstruct(tokens, true))
			}

			instructions, err := Parse(tokens)

			if DEBUG {
				PrintInstructions(instructions)
				fmt.Println()
			}

			if err != nil {
				fmt.Print(err, "\n")
			} else {
				fmt.Println(Compile(instructions))
			}
		}
	}
}
