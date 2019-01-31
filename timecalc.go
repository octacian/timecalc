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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>> ")
	text, _ := reader.ReadString('\n')
	tokens, err := Tokenize(text)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(Reconstruct(tokens))
	if os.Getenv("DEBUG") == "TRUE" {
		DEBUG = true
		fmt.Println("Debug mode enabled.")
	}
	}
}
