package tokenizer

import "strings"

// ExtractVariable ...
func ExtractVariable(input string) (string, string) {
	variable := ""
	if separatorIndex := strings.IndexRune(input, '='); separatorIndex != -1 {
		variable = strings.TrimSpace(input[:separatorIndex])
		input = input[separatorIndex+1:]
	}

	return variable, input
}
