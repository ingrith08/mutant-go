package validator

import "testing"

func TestValidateSequences(t *testing.T) {
	row := "CCCCTA"
	expected := true

	got := ValidateSequences(row)
	if got != expected {
		t.Errorf("ValidateSequences failed")
	}
}

func TestValidateNitrogenousBase(t *testing.T) {
	row := "CCCCTA"
	nitrogenousBase := "C"
	expected := true

	got := ValidateNitrogenousBase(row, nitrogenousBase)
	if got != expected {
		t.Errorf("ValidateNitrogenousBase failed")
	}
}
