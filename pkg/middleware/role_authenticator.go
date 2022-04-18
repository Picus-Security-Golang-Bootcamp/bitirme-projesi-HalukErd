package mw

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var m = map[string][]string{
	"/signup_POST": {"admin"},
}

func AuthenticateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxValue, exists := c.Get("roles")
		endpointKey := fmt.Sprintf("%s_%s", c.Request.RequestURI, c.Request.Method)
		zap.L().Info(endpointKey)
		if !exists {
			c.JSON(http.StatusForbidden, "There is no permission to create user.")
		}

		roles := ctxValue.([]string)
		for _, role := range roles {
			zap.L().Info(role)
		}
		c.Next()
		c.Abort()
		return
	}
}
