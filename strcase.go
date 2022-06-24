// Package strcase converts strings to various cases.
package strcase

import "strings"

// ToDelimited converts string to delimited.case in case delim = '.'.
func ToDelimited(s string, delim rune) string {
	return strings.Join(splitToWords(s), string(delim))
}

// ToSnake converts string to snake_case.
func ToSnake(s string) string {
	return ToDelimited(s, '_')
}

// ToKebab converts string to kebab-case.
func ToKebab(s string) string {
	return ToDelimited(s, '-')
}

// ToCamel converts string to camelCase.
func ToCamel(s string) string {
	words := splitToWords(s)

	if len(words) == 1 {
		return words[0]
	}

	n := 0
	for i := 0; i < len(words); i++ {
		n += len(words[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(words[0])
	for _, word := range words[1:] {
		b.WriteString(camelizeWord(word))
	}
	return b.String()
}

// ToUpperCamel converts string to UpperCamelCase.
func ToUpperCamel(s string) string {
	words := splitToWords(s)

	if len(words) == 1 {
		return camelizeWord(words[0])
	}

	n := 0
	for i := 0; i < len(words); i++ {
		n += len(words[i])
	}

	var b strings.Builder
	b.Grow(n)
	for _, word := range words {
		b.WriteString(camelizeWord(word))
	}
	return b.String()
}
