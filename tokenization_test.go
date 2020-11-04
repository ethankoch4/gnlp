package gnlp_test

import (
	"testing"

	"github.com/ethankoch4/gnlp"
)

func TestWhiteSpaceTokenizer(t *testing.T) {
	text := "My name is Ethan Koch."
	predicted := gnlp.WhiteSpaceTokenizer().tokenize(text)
	
	correct := ["My", "name", "is", "Ethan", "Koch"]
	if predicted == correct {
		t.Log("Correct!")
	} else {
		t.FailF("Wrong. Expected: %v but got: %v", correct, predicted)
	}
}