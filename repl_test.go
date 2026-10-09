package main

import ("testing")

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "HELLO WoRlD ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  GREAT DAY EVERYBODY AND GOOD NIGHT ",
			expected: []string{"great", "day", "everybody", "and", "good", "night"},
		},
		{
			input: " Save FIGHT drag captURE ",
			expected: []string{"save", "fight", "drag", "capture"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length should be %d not %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %q and got %q", expectedWord, word)
				t.Fatal()
			}
		}
	}
}
