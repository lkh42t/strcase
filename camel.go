package strcase

import (
	"strings"
	"unicode"
)

// ToCamel converts string to camelCase.
func ToCamel(s string) string {
	words := Parse(s)
	for i := 1; i < len(words); i++ {
		words[i] = camelize(words[i])
	}
	return strings.Join(words, "")
}

// ToUpperCamel converts string to UpperCamelCase.
func ToUpperCamel(s string) string {
	return mapJoin(Parse(s), "", camelize)
}

func camelize(word string) string {
	switch len(word) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(word)
	}

	rs := []rune(word)

	var b strings.Builder
	b.Grow(len(word))
	b.WriteRune(unicode.ToUpper(rs[0]))
	for _, r := range rs[1:] {
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
