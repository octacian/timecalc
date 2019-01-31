package main

import (
	"fmt"
	"math"
	"strconv"
)

// Factor represents a single component of a mathematical statement
type Factor interface {
	Raw() float64
	SetRaw(raw float64)
	String() string

	Add(factor Factor) error
	Subtract(factor Factor) error
	//Multiply(factor Factor) error
	//Divide(factor Factor) error
}

// NumberRegex represents any normal number in a mathematical statement
type Number struct {
	value float64
}

// String return the contents of the number as a string
func (n Number) String() string {
	if math.Floor(n.value) == n.value {
		return fmt.Sprintf("%v", n.value)
	}

	str := fmt.Sprintf("%f", n.value)

	// Trim trailing zeros
	for {
		if str[len(str)-1] != '0' {
			break
		}

		str = str[0 : len(str)-1]
	}

	return str
}

// Raw returns the value of a NumberRegex as a 64-bit float for processing
func (n Number) Raw() float64 {
	return n.value
}

// SetRaw sets the value of the number given a raw float64
func (n *Number) SetRaw(raw float64) {
	n.value = raw
}

// Add takes another factor and adds it to the current
func (n *Number) Add(factor Factor) error {
	raw := n.Raw()
	factorRaw := factor.Raw()

	n.SetRaw(raw + factorRaw)

	return nil
}

// Subtract takes another factor and subtracts it from the current
func (n *Number) Subtract(factor Factor) error {
	raw := n.Raw()
	factorRaw := factor.Raw()

	n.SetRaw(raw - factorRaw)

	return nil
}

// TimeRegex represents an arbitrary time down to the second
type Time struct {
	Hour   int
	Minute int
	Second int
}

// Raw returns the value of a Time as a 64-bit float for processing
func (t Time) Raw() float64 {
	return float64(t.Hour * 3600 + t.Minute * 60 + t.Second)
}

// SetRaw sets the value of the Time given a float64
func (t *Time) SetRaw(raw float64) {
	rawInt := int(raw)
	PrintDebug(fmt.Sprintf("SetRaw to %f (simplified to %d)\n", raw, rawInt))
	t.Hour = rawInt / 3600
	t.Minute = rawInt % 3600 / 60
	t.Second = rawInt % 3600 % 60
	PrintDebug(fmt.Sprintf("|- %s\n", t))
}

func fieldToString(field int) string {
	converted := fmt.Sprintf("%v", field)
	if converted == "0" {
		return "00"
	}

	if len(converted) == 1 {
		return "0" + converted
	}

	return converted
}

// String returns a string representation of the TimeRegex
func (t Time) String() string {
	//if t.Second == 0 {
	//	return fmt.Sprintf("%v:%v", fieldToString(t.Hour), fieldToString(t.Minute))
	//}

	return fmt.Sprintf("%v:%v:%v", fieldToString(t.Hour), fieldToString(t.Minute), fieldToString(t.Second))
}

// Add takes another factor and adds it to the current
func (t *Time) Add(factor Factor) error {
	raw := t.Raw()
	factorRaw := factor.Raw()

	t.SetRaw(raw + factorRaw)

	return nil
}

// Subtract takes another factor and subtracts it from the current
func (t *Time) Subtract(factor Factor) error {
	raw := t.Raw()
	factorRaw := factor.Raw()

	t.SetRaw(raw - factorRaw)

	return nil
}

// NewTime takes a string and returns a time or an error
func NewTime(time string) (Time, error) {
	matches := TimeRegex.FindStringSubmatch(time)
	// Replace blank strings with "0"
	for key, match := range matches {
		if match == "" {
			matches[key] = "0"
		}
	}
	hour, err := strconv.ParseInt(matches[1], 10, 0)
	if err != nil {
		return Time{}, err
	}
	minute, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		return Time{}, err
	}
	second, err := strconv.ParseInt(matches[3], 10, 0)
	if err != nil {
		return Time{}, nil
	}

	return Time{Hour: int(hour), Minute: int(minute), Second: int(second)}, nil
}

// Instruction represents a part of a statement in which a particular
// Operation is to be carried out upon two terms, Left and Right. If one
// of these terms is nil, no operation is performed. If the Operation
// is undefined, it defaults to multiplication. Each of Left and Right
// can be another Instruction set, a NumberRegex, or a TimeRegex.
type Instruction struct {
	Operation string
	Value     Factor
}

// Parse takes an array of tokens and returns an array of instructions
func Parse(tokens []*Token) ([]*Instruction, error) {
	instructions := make([]*Instruction, 0)
	current := &Instruction{Operation: "add"}

	for _, v := range tokens {
		switch v.Type {
		case "whitespace":
			continue
		case "operator":
			switch v.Value {
			case "+":
				current.Operation = "add"
			case "-":
				current.Operation = "subtract"
			}
		case "number":
			val, err := strconv.ParseFloat(v.Value, 64)
			if err != nil {
				panic(err)
			}
			current.Value = &Number{val}
		case "time":
			time, err := NewTime(v.Value)
			if err != nil {
				panic(err)
			}
			current.Value = &time
		}

		// if Value has been set, add to output
		if current.Value != nil {
			instructions = append(instructions, current)
			current = &Instruction{Operation: "add"}
		}
	}

	return instructions, nil
}

// PrintInstructions takes an array of instructions and prints them in plain text
func PrintInstructions(instructions []*Instruction) {
	for _, v := range instructions {
		fmt.Printf("%s: %+v\n\n", v.Operation, v.Value)
	}
}
