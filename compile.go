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

// Compile takes an array of instructions and processes them returning an output Factor
func Compile(instructions []*Instruction) Factor {
	var out Factor
	out = &Number{}

	// Loop through instructions and apply them to out
	for _, instruction := range instructions {
		value := instruction.Value
		// if value is a list of sub-instructions, compile those first
		if _, isList := value.([]*Instruction); isList {
			value = Compile(value.([]*Instruction))
		}

		switch instruction.Operation {
		case "+":
			left := out.Raw()
			right := value.(Factor).Raw()
			out.SetRaw(left + right)
		case "-":
			left := out.Raw()
			right := value.(Factor).Raw()
			out.SetRaw(left - right)
		case "*":
			left := out.Raw()
			right := value.(Factor).Raw()
			out.SetRaw(left * right)
		case "/":
			left := out.Raw()
			right := value.(Factor).Raw()
			out.SetRaw(left / right)
		}

		// if value has a greater weight than current output type, convert to this type
		if getWeight(value.(Factor)) > getWeight(out) {
			outRaw := out.Raw()
			out = value.(Factor)
			out.SetRaw(outRaw)
		}
	}

	return out
}
