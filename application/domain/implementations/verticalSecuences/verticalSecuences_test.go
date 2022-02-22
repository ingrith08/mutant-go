package verticalSecuences

import "testing"

func TestVerticalSecuences(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	expected := 1

	got := GetSecuences(dna)
	if got != expected {
		t.Errorf("GetSecuences Vertical failed, got: %d, want: %d.", got, expected)
	}
}

func TestGetColumn(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	expected := "ACTACT"
	got := getColumn(0, dna)
	if got != expected {
		t.Errorf("GetSecuences getColumn failed")
	}
}
