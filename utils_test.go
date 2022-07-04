package strcase

import "testing"

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
