package middleware

import (
	"fmt"
	"strings"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	// Authorizationヘッダからトークンを取り出す
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader != "" {
		ary := strings.Split(authorizationHeader, " ")
		if len(ary) == 2 {
			if ary[0] == "Bearer" {
				fmt.Println(ary[1])
			}
		}
	}
	c.Next()
}
