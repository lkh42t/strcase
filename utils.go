package strcase

import "strings"

func mapJoin(words []string, sep string, convert func(string) string) string {
	slice := make([]string, 0, len(words))
	for _, word := range words {
		slice = append(slice, convert(word))
	}
	return strings.Join(slice, sep)
}
