package time

import (
	"fmt"
	"regexp"
)

var OperatorRegex = regexp.MustCompile(`^(\+|\-|\*|\/|\%)$`)
var TimeRegex = regexp.MustCompile(`^([\d][\d]?)?:([\d][\d])?:?(?::([\d][\d]))?(?:\.([\d]+))?$`)
var NumberRegex = regexp.MustCompile(`^((\d+)?\.?)\d+$`)
var WhitespaceRegex = regexp.MustCompile(`^([\s])$`)

// Token represents a part of some user input
type Token struct {
	Type  string
	Value string
}

// Tokenize a string, returning an array of tokens
func Tokenize(str string) ([]*Token, error) {
	tokens := make([]*Token, 0)
	last := ""

	openGroupCount := 0
	closeGroupCount := 0

	for i, v := range str {
		c := string(v)
		var next string

		if i+1 < len(str) {
			next = string(str[i+1])
		}

		if OperatorRegex.MatchString(c) {
			tokens = append(tokens, &Token{"operator", c})

			if OperatorRegex.MatchString(next) {
				return nil, fmt.Errorf("tokenize: unexpected operator")
			}
		} else if NumberRegex.MatchString(c) {
			if NumberRegex.MatchString(next) || next == ":" || next == "." {
				last += c
				continue
			} else if last != "" {
				last += c
			}

			if last == "" {
				tokens = append(tokens, &Token{"number", c})
			}
		} else if c == ":" || c == "." {
			last += c
		} else if c == "(" || c == "[" {
			tokens = append(tokens, &Token{"groupOpen", c})
			openGroupCount++
		} else if c == ")" || c == "]" {
			tokens = append(tokens, &Token{"groupClose", c})
			closeGroupCount++
		} else if WhitespaceRegex.MatchString(c) {
			tokens = append(tokens, &Token{"whitespace", c})
		} else {
			return tokens, fmt.Errorf("tokenize: invalid charactor: %v", c)
		}

		if last != "" && !(NumberRegex.MatchString(next) || next == ":" || next == ".") {
			if NumberRegex.MatchString(last) {
				tokens = append(tokens, &Token{"number", last})
			} else if TimeRegex.MatchString(last) {
				tokens = append(tokens, &Token{"time", last})
			} else {
				return nil, fmt.Errorf("tokenize: invalid time or number")
			}

			last = ""
		}
	}

	// if an unequal number of group openings and closings were found, return error
	if openGroupCount != closeGroupCount {
		return nil, fmt.Errorf("tokenize: uneqal number of group openings and closings found (%d opened, "+
			"%d closed)", openGroupCount, closeGroupCount)
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
