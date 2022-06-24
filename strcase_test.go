package strcase

import "testing"

func TestToDelimited(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo", want: "foo"},
		{in: "foo bar", want: "foo.bar"},
	}

	for _, tt := range cases {
		got := ToDelimited(tt.in, '.')
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}

func TestToSnake(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo", want: "foo"},
		{in: "foo bar", want: "foo_bar"},
	}

	for _, tt := range cases {
		got := ToSnake(tt.in)
		if got != tt.want {
			t.Fatalf("wants %v, got %v", tt.want, got)
		}
	}
}

func TestToKebab(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "foo", want: "foo"},
		{in: "foo bar", want: "foo-bar"},
	}

	for _, tt := range cases {
		got := ToKebab(tt.in)
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
