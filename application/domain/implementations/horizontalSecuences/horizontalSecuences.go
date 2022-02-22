package horizontalSecuences

import (
	"test.com/mutant-go/application/domain/implementations/validator"
)

func GetSecuences(Dna []string) int {
	var sequences int = 0
	for _, row := range Dna {
		if validator.ValidateSequences(row) {
			sequences++
			if sequences > 1 {
				break
			}
		}
	}
	return sequences
}
