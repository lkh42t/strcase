package strcase

// ToSnake is an alias for ToLowerSnake.
var ToSnake = ToLowerSnake

// ToLowerSnake converts string to lower_snake_case.
func ToLowerSnake(s string) string {
	return ToLowerDelimited(s, "_")
}

// ToUpperSnake converts string to UPPER_SNAKE_CASE.
func ToUpperSnake(s string) string {
	return ToUpperDelimited(s, "_")
}
