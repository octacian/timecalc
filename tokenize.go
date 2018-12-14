package main

import (
	"fmt"
	"regexp"
)

// Token represents a part of some user input
type Token struct {
	Type  string
	Value string
}

// Tokenize a string, returning an array of tokens
func Tokenize(str string) ([]*Token, error) {
	tokens := make([]*Token, 0)
	last := ""

	for i, v := range str {
		c := string(v)

		operator := regexp.MustCompile(`^(\+|\-|\*|\/)$`)
		//time := regexp.MustCompile(`^([\d][\d]?):([\d][\d])(?::[\d][\d])?$`)
		//number := regexp.MustCompile(`^([\d]+(?:\.[\d]+)?)$`)
		number := regexp.MustCompile(`^([0-9])$`)
		whitespace := regexp.MustCompile(`^([\s])$`)

		if operator.MatchString(c) {
			tokens = append(tokens, &Token{"operator", c})
		} else if number.MatchString(c) {
			if number.MatchString(string(str[i+1])) {
				last += c
				continue
			} else if last != "" {
				last += c
			}

			if last == "" {
				tokens = append(tokens, &Token{"number", c})
			}
		} else if c == ":" || c == "." {
			tokens = append(tokens, &Token{"divider", c})
		} else if c == "(" || c == "[" {
			tokens = append(tokens, &Token{"groupOpen", c})
		} else if c == ")" || c == "]" {
			tokens = append(tokens, &Token{"groupClose", c})
		} else if whitespace.MatchString(c) {
			tokens = append(tokens, &Token{"whitespace", c})
		} else {
			return tokens, fmt.Errorf("tokenize: invalid charactor: %v", c)
		}

		if last != "" {
			tokens = append(tokens, &Token{"number", last})
			last = ""
		}
	}

	return tokens, nil
}

// Reconstruct a string from a token list
func Reconstruct(tokens []*Token) string {
	out := ""
	for _, token := range tokens {
		out += token.Value
	}

	return out
}
