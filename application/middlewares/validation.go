package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/validator.v2"
	"test.com/mutant-go/application/domain/types"
)

func Validation(ctx *gin.Context) {
	var body types.MutantBody
	response := types.MutantResponse{Status: http.StatusInternalServerError}
	err := ctx.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response)
	}

	ValidateStruct(ctx, &body)
	ValidateLength(ctx, &body)
	ValidateLetters(ctx, &body)
	ctx.Next()
}

func ValidateStruct(ctx *gin.Context, body *types.MutantBody) {
	err := validator.Validate(body)
	fmt.Println("err", err)
	if err != nil {
		json.Marshal(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "sólo se permiten secuencias desde 4x4 hasta 10X10")
	}
}

func ValidateLength(ctx *gin.Context, body *types.MutantBody) {
	var err error
	numberOfSequences := len(body.Dna)
	for _, nitrogenBase := range body.Dna {
		if numberOfSequences != len(nitrogenBase) {
			err = errors.New("la matriz debe ser simetrica (NxN)")
		}
	}
	if err != nil {
		json.Marshal(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
}

func ValidateLetters(ctx *gin.Context, body *types.MutantBody) {
	var err error
	for _, base := range body.Dna {
		for _, Letter := range strings.Split(base, "") {
			if !strings.Contains("ATGC", Letter) {
				err = errors.New("letra de base nitrogenada inválida")
			}
		}
	}
	if err != nil {
		json.Marshal(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
}
