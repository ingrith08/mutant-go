package models

type Mutant struct {
	IsMutant bool     `json:"isMutant"`
	Dna      []string `json:"dna"`
	ID       string   `bson:"_id" json:"id"`
}
