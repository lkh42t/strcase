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

func TestToCamel(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo", want: "foo"},
		{in: "foo bar", want: "fooBar"},
	}

	for _, tt := range cases {
		got := ToCamel(tt.in)
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}

func TestToUpperCamel(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo", want: "Foo"},
		{in: "foo bar", want: "FooBar"},
	}

	for _, tt := range cases {
		got := ToUpperCamel(tt.in)
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}
