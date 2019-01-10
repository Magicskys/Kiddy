package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"Kiddy/setting"
	"net/http"
)
type CustomClaims struct {
	ID int `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func MyMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth,err:=c.Request.Cookie("PHPSESSION");if err==nil{
			token, err := jwt.Parse(auth.Value,
				func(token *jwt.Token) (interface{}, error) {
					return []byte(setting.JwtSecret), nil
				})
			if err == nil {
				if token.Valid {
					c.Next()
					return
				} else {
					c.String(http.StatusUnauthorized, "Token is not valid")
					c.Abort()
					return
				}
			}
		}

		c.String(http.StatusUnauthorized, "Unauthorized access to this resource")
		c.Abort()
		return
	}
}