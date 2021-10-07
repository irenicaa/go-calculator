package tokenizer

import "strings"

// RemoveComment ...
func RemoveComment(input string) string {
	if separatorIndex := strings.Index(input, "//"); separatorIndex != -1 {
		input = input[:separatorIndex]
	}

	return input
}
