package mw

import (
	"BasketProjectGolang/pkg/config"
	jwtHelper "BasketProjectGolang/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, cfg.JWTConfig.TokenPrefix) {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
			c.Next()
			c.Abort()
			return
		}
		tokenStr := strings.Replace(authorizationHeader, cfg.JWTConfig.TokenPrefix, "", 1)
		decodedClaims := jwtHelper.VerifyToken(tokenStr, cfg.JWTConfig.SecretKey)
		if decodedClaims == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		}
		c.Set("roles", decodedClaims.Roles)
		c.Next()
		c.Abort()
		return
	}
}
