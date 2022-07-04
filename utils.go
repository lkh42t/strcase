package strcase

import (
	"strings"
	"unicode"
)

func mapJoin(words []string, sep string, convert func(string) string) string {
	for i := 0; i < len(words); i++ {
		words[i] = convert(words[i])
	}
	return strings.Join(words, sep)
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
