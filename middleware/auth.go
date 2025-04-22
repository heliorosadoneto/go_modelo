package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user_id") == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Não autorizado"})
			c.Abort()
			return
		}
		c.Next()
	}
}
