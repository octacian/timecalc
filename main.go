package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/peterh/liner"
)

// DEBUG defines whether extra verbose information is printed
var DEBUG bool

// PrintDebug prints a string only if debug mode is enabled
func PrintDebug(str string) {
	if DEBUG {
		fmt.Print(str)
	}
}

var OperatorRegex = regexp.MustCompile(`^(\+|\-|\*|\/|\%)$`)
var TimeRegex = regexp.MustCompile(`^([\d][\d]?)?:([\d][\d])?:?(?::([\d][\d]))?(?:\.([\d]+))?$`)
var NumberRegex = regexp.MustCompile(`^((\d+)?\.?)\d+$`)
var WhitespaceRegex = regexp.MustCompile(`^([\s])$`)

var history_fn = filepath.Join(os.TempDir(), "timecalc.tmp")

func main() {
	if os.Getenv("DEBUG") == "TRUE" {
		DEBUG = true
		fmt.Println("Debug mode enabled.")
	}

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	if f, err := os.Open(history_fn); err != nil {
		line.ReadHistory(f)
		f.Close()
	}

	for {
		input, err := line.Prompt(">>> ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				return
			}
			log.Fatalf("error reading line: %s", err)
		}

		if input == "exit" {
			return
		}

		line.AppendHistory(input)

		tokens, err := Tokenize(input)
		if err != nil {
			fmt.Println(err)
		}
		if DEBUG {
			fmt.Println(Reconstruct(tokens, true))
		}

		instructions, err := Parse(tokens)

		if DEBUG {
			PrintInstructions(instructions)
			fmt.Println()
		}

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(Compile(instructions))
		}
	}
}
