package tokenizer

import "strings"

// ExtractVariable ...
func ExtractVariable(input string) (variable string, code string) {
	code = input
	if separatorIndex := strings.IndexRune(input, '='); separatorIndex != -1 {
		variable = strings.TrimSpace(input[:separatorIndex])
		code = input[separatorIndex+1:]
	}

	return variable, code
}
