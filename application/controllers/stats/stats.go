package stats

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/mutant-go/application/repositories_impl"
)

func GetStats(ctx *gin.Context) {
	var repo = &repositories_impl.MutantRepo
	response := repo.GetStats()
	ctx.JSON(http.StatusOK, response)
}
