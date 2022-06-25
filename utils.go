package strcase

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func parse(s string) []string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	var words []string

	start := 0
	end := 0
	runeWidth := 1
	consectiveUppers := 0
	for i, r := range s {
		if unicode.IsUpper(r) {
			if consectiveUppers == 0 && i != 0 {
				words = append(words, s[start:end])
				start = end
			}
			consectiveUppers++
			runeWidth = utf8.RuneLen(r)
			end += runeWidth
		} else if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if consectiveUppers > 1 {
				words = append(words, s[start:end-runeWidth])
				start = end - runeWidth
			}
			consectiveUppers = 0
			runeWidth = utf8.RuneLen(r)
			end += runeWidth
		} else {
			if end-start != 0 {
				words = append(words, s[start:end])
				consectiveUppers = 0
			}
			runeWidth = utf8.RuneLen(r)
			end += runeWidth
			start = end
		}
	}
	return append(words, s[start:end])
}

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
