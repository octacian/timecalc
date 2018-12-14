package main

import (
	"bufio"
	"fmt"
	"os"
)

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
