package strcase

import "testing"

func TestToLowerKebab(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo bar", want: "foo-bar"},
	}

	for _, tt := range cases {
		got := ToLowerKebab(tt.in)
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}

func TestToUpperKebab(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo bar", want: "FOO-BAR"},
	}

	for _, tt := range cases {
		got := ToUpperKebab(tt.in)
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}
