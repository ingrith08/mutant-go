package diagonalSecuences

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiagonalSecuences(t *testing.T) {
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
		t.Errorf("GetSecuences Diagonal failed, got: %d, want: %d.", got, expected)
	}
}

func TestGetSecuence(t *testing.T) {
	diagonalSequence := []string{
		"A",
		"CT",
		"TAG",
		"ATGC",
		"CGATG",
		"TCATGA",
		"CCAGC",
		"ACGT",
		"CTG",
		"TA",
		"C",
	}
	diagonalSequenceReverse := []string{
		"A",
		"GC",
		"CGT",
		"GTGG",
		"TGTGA",
		"AAAATC",
		"CTACT",
		"TGCC",
		"ACA",
		"CC",
		"T",
	}
	expected1 := 0
	expected2 := 1
	got1 := getSecuence(diagonalSequence)
	got2 := getSecuence(diagonalSequenceReverse)
	if got1 != expected1 {
		t.Errorf("GetSecuence diagonal failed, got: %d, want: %d.", got1, expected1)
	}
	if got2 != expected2 {
		t.Errorf("GetSecuence diagonal reverse failed, got: %d, want: %d.", got2, expected2)
	}
}

func TestGetDiagonalSequence(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC",
	}
	expected1 := []string{
		"A",
		"CT",
		"TAG",
		"ATGC",
		"CGATG",
		"TCATGA",
		"CCAGC",
		"ACGT",
		"CTG",
		"TA",
		"C",
	}
	expected2 := []string{
		"T",
		"CC",
		"ACA",
		"CCGT",
		"TCATC",
		"CTAAAA",
		"AGTGT",
		"GGTG",
		"TGC",
		"CG",
	}
	got1 := getDiagonalSequence(dna, false)
	got2 := getDiagonalSequence(dna, true)
	if !reflect.DeepEqual(got1, expected1) {
		fmt.Println("got1", got1)
		t.Errorf("GetDiagonalSequence failed")
	}
	if !reflect.DeepEqual(got2, expected2) {
		fmt.Println("got1", got2)
		t.Errorf("GetDiagonalSequence reverse failed")
	}
}
