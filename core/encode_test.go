package core

import "testing"

func TestPadSepString(t *testing.T) {
	cases := []struct {
		s           string
		insertEvery int
		sep         string
		expected    string
	}{
		{"tooshort", 12, "-", "tooshort"},
		{"helloworld", 5, "-", "hello-world"},
	}

	for _, c := range cases {
		if actual := padSepString(c.s, c.insertEvery, c.sep); actual != c.expected {
			t.Logf("should be '%s' got '%s'", c.expected, actual)
			t.Fail()
		}
	}
}
