package ast

import (
	"regexp"
	"strconv"
	"strings"
)

// SimpleStringValue checks if a value is a string containing only basic symbols.
// in sync with [a-zA-Z0-9-._:/+;~@<>]+ in PEG (with ^ and $ around)
var SimpleStringValue = regexp.MustCompile("^[a-zA-Z0-9-._:/+;~@<>*]+$")

func quoteStringIfNeeded(input string) string {
	if _, err := strconv.Atoi(input); err == nil {
		return "'" + input + "'"
	}
	if _, err := strconv.ParseFloat(input, 64); err == nil {
		return "'" + input + "'"
	}
	if SimpleStringValue.MatchString(input) {
		return input
	}
	return Quote(input)
}

// Quote a string.
func Quote(str string) string {
	if strings.ContainsRune(str, '\'') {
		return "\"" + str + "\""
	}
	return "'" + str + "'"
}

func isQuoted(str string) bool {
	return len(str) > 1 && ((str[0] == '\'' && str[len(str)-1] == '\'') ||
		(str[0] == '"' && str[len(str)-1] == '"'))
}
