package horizontalSecuences

import "testing"

func TestHorizontalSecuences(t *testing.T) {
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
		t.Errorf("GetSecuences Horizontal failed, got: %d, want: %d.", got, expected)
	}
}
