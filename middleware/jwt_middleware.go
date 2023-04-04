package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/VinncentWong/DiverBE/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.Abort()
			util.SendResponse(c, http.StatusForbidden, "wrong token type", false, nil)
			return
		}
		token := header[7:]
		jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			c.Abort()
			util.SendResponse(c, http.StatusForbidden, err.Error(), false, nil)
			return
		}
		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if jwtToken.Valid && ok {
			c.Set("username", claims["username"])
			c.Next()
		} else {
			c.Abort()
			util.SendResponse(c, http.StatusForbidden, err.Error(), false, nil)
			return
		}
	}
}
