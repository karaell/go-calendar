package events

import "testing"

func TestIsValidTitle(t *testing.T) {
	title := "Test 1"
	isValid := isValidTitle(title)
	if !isValid {
		t.Errorf("Valid title: expected true, but got false")
	}

	title = "Test incorrect )"
	isValid = isValidTitle(title)
	if isValid {
		t.Errorf("Invalid symbol: expected false, but got true")
	}

	title = "T"
	isValid = isValidTitle(title)
	if isValid {
		t.Errorf("Invalid length: expected false, but got true")
	}

	title = ""
	isValid = isValidTitle(title)
	if isValid {
		t.Errorf("Empty title: expected false, but got true")
	}
}
