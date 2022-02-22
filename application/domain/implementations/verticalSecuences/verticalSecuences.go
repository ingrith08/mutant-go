package verticalSecuences

import (
	"test.com/mutant-go/application/domain/implementations/validator"
)

func GetSecuences(Dna []string) int {
	var sequences int = 0
	row := ""
	for index := range Dna {
		row = getColumn(index, Dna)
		if validator.ValidateSequences(row) {
			sequences++
			if sequences > 1 {
				break
			}
		}
	}
	return sequences
}

func getColumn(index int, Dna []string) string {
	column := ""
	for _, row := range Dna {
		column += string(row[index])
	}
	return column
}
