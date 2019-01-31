package main

// getWeight takes a Factor and returns its importance
func getWeight(factor Factor) int {
	switch factor.(type) {
	case *Number:
		return 0
	case *Time:
		return 1
	}

	return 0
}

// Compile takes an array of instructions and processes them returning a string
// for printing to the console.
func Compile(instructions []*Instruction) string {
	var out Factor
	out = &Number{}

	// Loop through instructions and apply them to out
	for _, instruction := range instructions {
		switch instruction.Operation {
		case "+":
			left := out.Raw()
			right := instruction.Value.Raw()
			out.SetRaw(left + right)
		case "-":
			left := out.Raw()
			right := instruction.Value.Raw()
			out.SetRaw(left - right)
		case "*":
			left := out.Raw()
			right := instruction.Value.Raw()
			out.SetRaw(left * right)
		case "/":
			left := out.Raw()
			right := instruction.Value.Raw()
			out.SetRaw(left / right)
		}

		// if value has a greater weight than current output type, convert to this type
		if getWeight(instruction.Value) > getWeight(out) {
			outRaw := out.Raw()
			out = instruction.Value
			out.SetRaw(outRaw)
		}
	}

	return out.String()
}
