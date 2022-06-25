// Package strcase converts strings to various cases.
package strcase

import "strings"

var (
	// ToDelimited is an alias for ToLowerDelimited.
	ToDelimited = ToLowerDelimited
	// ToSnake is an alias for ToLowerSnake.
	ToSnake = ToLowerSnake
	// ToKebab is an alias for ToLowerKebab.
	ToKebab = ToLowerKebab
)

// ToLowerDelimited converts string to lower.delimited.case in case `sep = "."`.
func ToLowerDelimited(s string, sep string) string {
	return mapJoin(s, sep, strings.ToLower)
}

// ToUpperDelimited converts string to UPPER.DELIMITED.CASE in case `sep = "."`.
func ToUpperDelimited(s string, sep string) string {
	return mapJoin(s, sep, strings.ToUpper)
}

// ToLowerSnake converts string to lower_snake_case.
func ToLowerSnake(s string) string {
	return ToLowerDelimited(s, "_")
}

// ToUpperSnake converts string to UPPER_SNAKE_CASE.
func ToUpperSnake(s string) string {
	return ToUpperDelimited(s, "_")
}

// ToLowerKebab converts string to lower-kebab-case.
func ToLowerKebab(s string) string {
	return ToLowerDelimited(s, "-")
}

// ToUpperKebab converts string to UPPER-KEBAB-CASE.
func ToUpperKebab(s string) string {
	return ToUpperDelimited(s, "-")
}

// ToCamel converts string to camelCase.
func ToCamel(s string) string {
	words := parse(s)

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
	words := parse(s)

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
