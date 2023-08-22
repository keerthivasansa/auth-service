package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthToken(c *gin.Context) string {
	authH := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authH, "Bearer ")
	return token
}
