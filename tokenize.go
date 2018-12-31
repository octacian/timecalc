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
		var next string

		if i+1 < len(str) {
			next = string(str[i+1])
		}

		operator := regexp.MustCompile(`^(\+|\-|\*|\/)$`)
		time := regexp.MustCompile(`^([\d][\d]?)?:([\d][\d])?:?(?::([\d][\d]))?$`)
		number := regexp.MustCompile(`^((\d+)?\.?)\d+$`)
		whitespace := regexp.MustCompile(`^([\s])$`)

		if operator.MatchString(c) {
			tokens = append(tokens, &Token{"operator", c})

			if operator.MatchString(next) {
				return nil, fmt.Errorf("invalid token set")
			}
		} else if number.MatchString(c) {
			if number.MatchString(next) || next == ":" || next == "." {
				last += c
				continue
			} else if last != "" {
				last += c
			}

			if last == "" {
				tokens = append(tokens, &Token{"number", c})
				fmt.Println("Appending number by itself")
			}
		} else if c == ":" || c == "." {
			last += c
		} else if c == "(" || c == "[" {
			tokens = append(tokens, &Token{"groupOpen", c})

			if next == "(" || next == "[" {
				return nil, fmt.Errorf("invalid token set")
			}
		} else if c == ")" || c == "]" {
			tokens = append(tokens, &Token{"groupClose", c})

			if next == ")" || next == "]" {
				return nil, fmt.Errorf("invalid token set")
			}
		} else if whitespace.MatchString(c) {
			tokens = append(tokens, &Token{"whitespace", c})
		} else {
			return tokens, fmt.Errorf("tokenize: invalid charactor: %v", c)
		}

		if last != "" && !(number.MatchString(next) || next == ":" || next == ".") {
			if number.MatchString(last) {
				tokens = append(tokens, &Token{"number", last})
			} else if time.MatchString(last) {
				tokens = append(tokens, &Token{"time", last})
			} else {
				return nil, fmt.Errorf("invalid token set (in final append)")
			}

			last = ""
		}
	}

	return tokens, nil
}

// Reconstruct a string from a token list
func Reconstruct(tokens []*Token, verbose bool) string {
	out := ""
	for _, token := range tokens {
		out += token.Value

		if verbose && token.Type != "whitespace" {
			out += " (" + token.Type + ")"
		}
	}

	return out
}
