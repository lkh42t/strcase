package strcase

// ToKebab is an alias for ToLowerKebab.
var ToKebab = ToLowerKebab

// ToLowerKebab converts string to lower-kebab-case.
func ToLowerKebab(s string) string {
	return ToLowerDelimited(s, "-")
}

// ToUpperKebab converts string to UPPER-KEBAB-CASE.
func ToUpperKebab(s string) string {
	return ToUpperDelimited(s, "-")
}
