package strcase

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{in: "", want: []string{""}},
		{in: "foo", want: []string{"foo"}},
		{in: "fooBar", want: []string{"foo", "Bar"}},
		{in: "Foo-bar", want: []string{"Foo", "bar"}},
		{in: "userID", want: []string{"user", "ID"}},
		{in: "foo-_bar", want: []string{"foo", "bar"}},
		{in: "FOOBar", want: []string{"FOO", "Bar"}},
		{in: "Foo123bar", want: []string{"Foo", "123", "bar"}},
		{in: "föo", want: []string{"föo"}},
	}

	for _, tt := range cases {
		got := Parse(tt.in)
		if len(got) != len(tt.want) {
			t.Fatalf("Parse(%v) => %v, want %v", tt.in, got, tt.want)
		}
		for i, v := range got {
			if v != tt.want[i] {
				t.Fatalf("Parse(%v) => %v, want %v", tt.in, got, tt.want)
			}
		}
	}
}
