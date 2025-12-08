package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
	input    string
	expected []string
	}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input: " Charizard Bulbasaur Squirtle",
		expected: []string{"charizard", "bulbasaur", "squirtle"},
	},
	// add more cases here
}

	for _,c := range cases {

		actual:=cleanInput(c.input)

		for i:=range actual {
			
			word := actual[i]
			expectedWord:=c.expected[i]

			if word != expectedWord {
				t.Errorf("%v and %v do not match in test case %v", word, expectedWord, i)
			}

		}

	}

}
