package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

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
	}
}
