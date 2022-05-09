package services

import (
	"net/http"

	"github.com/PAWProjetoFinal/backend/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*") // nao podemos mandar assim, mais tarde alterar isto, nao podemos permitir tudo

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func AuthorizationRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !ValidateTokenJWT(c) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
			c.Abort()
		} else {
			var tokenInput, _, _ = getAuthorizationToken(c)
			token, err := jwt.ParseWithClaims(tokenInput, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
				c.Abort()
				return
			}

			if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
				//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
				c.Set("username", claims.Username)
			}
			OpenDatabase()

			// before request
			c.Next()
		}
	}
}
