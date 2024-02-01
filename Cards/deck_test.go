package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDeckFileNewDeckFunction(t *testing.T) {

	d := NewDeckFunc()

	if len(d) != 4 {
		t.Errorf("Expected length 4, but got %d", len(d))
	}

	if d[0] != "Pankaj of  Jaipur" {
		t.Errorf("Expected d[0] value is Pankaj , But got %s ", d[0])
	}
}

func TestNewDecFileCreateChec(t *testing.T) {
	os.Remove("_DeckFile")

	deck := NewDeckFunc()
	deck.SavedFile("_DeckFile")
	loadDeckFile := ReadStringFile("_DeckFile")

	if len(loadDeckFile) != 13 {
		t.Errorf("Expectd File Length 13 but Expected %v", len(loadDeckFile))
	}

	os.Remove("_DeckFile")

}

// check even and odd nuber
func TestCheckEvenOrOddNumber(t *testing.T) {
	tests := []struct {
		input  ResultNuber
		output []string
	}{
		{
			input:  ResultNuber{2, 7, 4, 9, 6},
			output: []string{"2 is even", "7 is odd", "4 is even", "9 is odd", "6 is even"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %v", test.input), func(t *testing.T) {
			result := test.input.checkEvenOrOddNuber()

			if len(result) != len(test.output) {
				t.Errorf("Expected %d messages, but got %d", len(test.output), len(result))
				return
			}

			for i, expectedMessage := range test.output {
				if result[i] != expectedMessage {
					t.Errorf("Expected: %s, Got: %s", expectedMessage, result[i])
				}
			}
		})
	}
}
