package diagonalSecuences

import (
	"test.com/mutant-go/application/domain/implementations/validator"
)

func GetSecuences(Dna []string) int {
	diagonalSequence := getDiagonalSequence(Dna, false)
	diagonalSequenceReverse := getDiagonalSequence(Dna, true)

	return getSecuence(diagonalSequence) + getSecuence(diagonalSequenceReverse)
}

func getSecuence(diagonalSequence []string) int {
	var sequences int = 0
	for _, row := range diagonalSequence {
		if validator.ValidateSequences(row) {
			sequences++
			if sequences > 1 {
				break
			}
		}
	}
	return sequences
}

func getDiagonalSequence(Dna []string, isReverse bool) []string {
	numberSequences := len(Dna)
	var numberSequencesTwo = 2 * (numberSequences - 1)
	var diagonal []string

	for k := 0; k <= numberSequencesTwo; k++ {
		letters := ""
		for y := numberSequences - 1; y >= 0; y-- {
			var x int
			if isReverse {
				x = k - (numberSequences - y)
			} else {
				x = k - y
			}
			if x >= 0 && x < numberSequences {
				letters += string((Dna)[y][x])
			}
		}

		if len(letters) > 0 {
			diagonal = append(diagonal, letters)
		}
	}

	return diagonal
}
