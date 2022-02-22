package healtcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealtCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}
