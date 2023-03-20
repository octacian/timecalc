package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
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

func main() {
	if os.Getenv("DEBUG") == "TRUE" {
		DEBUG = true
		fmt.Println("Debug mode enabled.")
	}

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	for {
		reader := bufio.NewReader(os.Stdin)
		color.Set(color.FgCyan)
		fmt.Print(">>> ")

	termLoop:
		for {
			switch event := termbox.PollEvent(); event.Type {
			case termbox.EventKey:
				if event.Key == termbox.KeyArrowUp {

				} else if event.Key == termbox.KeyArrowDown {

				} else if event.Key == termbox.KeyEnter {
					panic("")
					//break termLoop
				}
			case termbox.EventError:
				panic(event.Err)
			case termbox.EventInterrupt:
				break termLoop
			}
		}

		text, _ := reader.ReadString('\n')
		color.Unset()
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
