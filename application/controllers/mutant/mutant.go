package mutant

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"test.com/mutant-go/application/domain/implementations/diagonalSecuences"
	"test.com/mutant-go/application/domain/implementations/horizontalSecuences"
	"test.com/mutant-go/application/domain/implementations/verticalSecuences"
	"test.com/mutant-go/application/domain/models"
	"test.com/mutant-go/application/domain/types"
	"test.com/mutant-go/application/repositories_impl"
)

func IsMutant(ctx *gin.Context) {
	var body types.MutantBody
	var repo = &repositories_impl.MutantRepo
	response := types.MutantResponse{Status: http.StatusOK}
	ctx.ShouldBindBodyWith(&body, binding.JSON)

	horizontalSecuences := horizontalSecuences.GetSecuences(body.Dna)
	verticalSecuences := verticalSecuences.GetSecuences(body.Dna)
	diagonalSecuences := diagonalSecuences.GetSecuences(body.Dna)

	totalSequencesFound := horizontalSecuences + verticalSecuences + diagonalSecuences
	isMutant := totalSequencesFound > 1
	dnaString := strings.Join(body.Dna, " ")
	mutant := models.Mutant{ID: getHash(dnaString), IsMutant: isMutant, Dna: body.Dna}
	repo.Insert(mutant)
	if isMutant {
		ctx.JSON(http.StatusOK, response)
	} else {
		response.Status = http.StatusForbidden
		ctx.JSON(http.StatusForbidden, response)
	}
}

func getHash(s string) string {
	hash := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", hash)
}
