package strcase

import "strings"

func mapJoin(words []string, sep string, convert func(string) string) string {
	for i := 0; i < len(words); i++ {
		words[i] = convert(words[i])
	}
	return strings.Join(words, sep)
}
