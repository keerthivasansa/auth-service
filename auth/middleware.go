package auth

import (
	"errors"
	"fmt"
	"kv/auth/utils"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

const CTX_USER_KEY = "user"

var UNAUTH_ROUTES = []string{
	"/ping/",
}

var ErrUnauthenticated = errors.New("you need to be authenticated to access this route")

func GetUser(c *gin.Context) User {
	usrI, _ := c.Get(CTX_USER_KEY)
	user, ok := usrI.(User)
	if !ok {
		c.AbortWithError(http.StatusUnauthorized, ErrUnauthenticated)
		return User{}
	}
	return user
}

func AuthMiddleware(c *gin.Context) {
	token := utils.GetAuthToken(c)
	fmt.Printf("Path: %s", c.Request.URL.Path)
	if !slices.Contains(UNAUTH_ROUTES, c.FullPath()) && token == "" {
		c.AbortWithError(http.StatusUnauthorized, ErrUnauthenticated)
		return
	}

}
