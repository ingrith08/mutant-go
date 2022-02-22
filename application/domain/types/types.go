package types

type MutantBody struct {
	Dna []string `json:"dna" validate:"min=4,max=10"`
}

type MutantResponse struct {
	Status int `json:"status"`
}

type Stats struct {
	Count_human_dna  int     `bson:"count_human_dna" json:"count_human_dna"`
	Count_mutant_dna int     `bson:"count_mutant_dna" json:"count_mutant_dna"`
	Ratio            float32 `bson:"ratio,truncate" json:"ratio"`
}
