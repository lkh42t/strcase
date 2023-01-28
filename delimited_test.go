package strcase

import "testing"

func TestToLowerDelimited(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo", want: "foo"},
		{in: "foo bar", want: "foo.bar"},
	}

	for _, tt := range cases {
		got := ToLowerDelimited(tt.in, ".")
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}

func TestToUpperDelimited(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo", want: "FOO"},
		{in: "foo bar", want: "FOO.BAR"},
	}

	for _, tt := range cases {
		got := ToUpperDelimited(tt.in, ".")
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}
