package strcase

import "testing"

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

func TestCamelize(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "", want: ""},
		{in: "a", want: "A"},
		{in: "foo", want: "Foo"},
	}

	for _, tt := range cases {
		got := camelize(tt.in)
		if got != tt.want {
			t.Fatalf("camelize(%v) => %v, want %v", tt.in, got, tt.want)
		}
	}
}
