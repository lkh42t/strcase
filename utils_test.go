package strcase

import "testing"

func TestMapJoin(t *testing.T) {
	cases := []struct {
		in  []string
		out string
	}{
		{in: []string{}, out: ""},
		{in: []string{"Foo"}, out: "Foo"},
		{in: []string{"Foo", "bAr"}, out: "Foo.bAr"},
	}

	for _, tt := range cases {
		got := mapJoin(tt.in, ".", func(s string) string { return s })
		if got != tt.out {
			t.Fatalf("mapJoin(%v, \".\", strings.ToLower) => %v, want %v", tt.in, got, tt.out)
		}
	}
}
