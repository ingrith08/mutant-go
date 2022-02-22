package routes

import (
	"test.com/mutant-go/application/controllers/healtcheck"
	"test.com/mutant-go/application/controllers/mutant"
	"test.com/mutant-go/application/controllers/stats"

	"github.com/gin-gonic/gin"
	"test.com/mutant-go/application/middlewares"
)

func Routes(routes *gin.Engine) {
	routes.GET("/healtcheck", healtcheck.HealtCheck)
	routes.POST("/mutant", middlewares.Validation, mutant.IsMutant)
	routes.GET("/stats", stats.GetStats)
}
