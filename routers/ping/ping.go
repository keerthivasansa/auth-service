package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(eng *gin.Engine) {
	pingRouter := eng.Group("/ping")

	pingRouter.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	})
}
