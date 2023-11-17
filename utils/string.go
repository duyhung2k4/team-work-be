package utils

import (
	"strings"
	"unicode"
)

func UppercaseFirstChar(str string) string {
	s := []rune(str)
	s[0] = unicode.ToUpper(s[0])
	return string(s)
}

func Lowercase(str string) string {
	return strings.ToLower(str)
}
