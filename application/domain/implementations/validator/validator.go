package validator

import "strings"

func ValidateSequences(row string) bool {
	return ValidateNitrogenousBase(row, "A") || ValidateNitrogenousBase(row, "T") ||
		ValidateNitrogenousBase(row, "G") || ValidateNitrogenousBase(row, "C")
}

func ValidateNitrogenousBase(row string, nitrogenousBase string) bool {
	equalSequence := strings.Repeat(nitrogenousBase, 4)
	return strings.Contains(row, equalSequence)
}
