package strcase

import "testing"

func TestSplitToWords(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{in: "", want: []string{""}},
		{in: "foo", want: []string{"foo"}},
		{in: "fooBar", want: []string{"foo", "bar"}},
		{in: "Foo-bar", want: []string{"foo", "bar"}},
		{in: "userID", want: []string{"user", "id"}},
		{in: "foo-_bar", want: []string{"foo", "bar"}},
		{in: "FOOBar", want: []string{"foo", "bar"}},
	}

	for _, tt := range cases {
		got := splitToWords(tt.in)
		if len(got) != len(tt.want) {
			t.Fatalf("splitToWords(%v) => %v, want %v", tt.in, got, tt.want)
		}
		for i, v := range got {
			if v != tt.want[i] {
				t.Fatalf("splitToWords(%v) => %v, want %v", tt.in, got, tt.want)
			}
		}
	}
}

func TestCamelizeWord(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "", want: ""},
		{in: "a", want: "A"},
		{in: "foo", want: "Foo"},
	}

	for _, tt := range cases {
		got := camelizeWord(tt.in)
		if got != tt.want {
			t.Fatalf("camelizeWord(%v) => %v, want %v", tt.in, got, tt.want)
		}
	}
}
