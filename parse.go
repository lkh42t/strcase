package strcase

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type token int

const (
	tokenLower token = iota
	tokenUpper
	tokenDigit
)

// Parse splits string to words (e.g. FooBar => [Foo Bar]).
//
// This drops all non letters or digits.
func Parse(s string) []string {
	// Trim non letters or digits
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	var words []string

	start := 0
	end := 0
	runeWidth := 1 // stores previous rune's width
	prevToken := tokenLower
	consectiveUppers := 0
	for i, r := range s {
		switch {
		case unicode.IsUpper(r):
			if prevToken != tokenUpper && i != 0 {
				words = append(words, s[start:end])
				start = end
			}
			consectiveUppers++
			prevToken = tokenUpper
			runeWidth = utf8.RuneLen(r)
			end += runeWidth
		case unicode.IsLower(r):
			if prevToken == tokenUpper && consectiveUppers > 1 {
				// e.g. FOOBar
				//      |  |`-current
				//      |  `end
				//      start
				// In this case, we want to get FOO.
				words = append(words, s[start:end-runeWidth])
				start = end - runeWidth
			} else if prevToken == tokenDigit {
				words = append(words, s[start:end])
				start = end
			}
			consectiveUppers = 0
			prevToken = tokenLower
			runeWidth = utf8.RuneLen(r)
			end += runeWidth
		case unicode.IsDigit(r):
			if prevToken != tokenDigit {
				words = append(words, s[start:end])
				start = end
			}
			consectiveUppers = 0
			prevToken = tokenDigit
			runeWidth = utf8.RuneLen(r)
			end += runeWidth
		default:
			if end-start != 0 {
				words = append(words, s[start:end])
				consectiveUppers = 0
				prevToken = tokenLower
			}
			runeWidth = utf8.RuneLen(r)
			end += runeWidth
			start = end
		}
	}
	return append(words, s[start:end])
}
