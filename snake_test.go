package strcase

import "testing"

func TestToLowerSnake(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo bar", want: "foo_bar"},
	}

	for _, tt := range cases {
		got := ToLowerSnake(tt.in)
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}

func TestToUpperSnake(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo bar", want: "FOO_BAR"},
	}

	for _, tt := range cases {
		got := ToUpperSnake(tt.in)
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}
