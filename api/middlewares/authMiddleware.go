package middlewares

import (
	"net/http"
	"strings"

	"github.com/MaheshBailwal/mscore/core"
	"github.com/MaheshBailwal/mscore/security"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if strings.Contains(strings.ToLower(c.Request.URL.Path), "/api/") {
			token := c.Request.Header.Get("Authorization")
			token = strings.ReplaceAll(token, "Bearer ", "")

			claims, err := security.ValidateToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Token Not found"})
				c.Abort()
				return
			}

			sc := core.NewServiceContext(claims["userId"].(string), c)

			c.Set("ServiceContext", sc)
		}

		c.Next()
	}

}
