package strcase

import "strings"

// ToDelimited is an alias for ToLowerDelimited.
var ToDelimited = ToLowerDelimited

// ToLowerDelimited converts string to lower.delimited.case in case `sep = "."`.
func ToLowerDelimited(s, sep string) string {
	return mapJoin(Parse(s), sep, strings.ToLower)
}

// ToUpperDelimited converts string to UPPER.DELIMITED.CASE in case `sep = "."`.
func ToUpperDelimited(s, sep string) string {
	return mapJoin(Parse(s), sep, strings.ToUpper)
}
