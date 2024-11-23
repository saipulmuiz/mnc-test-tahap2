package middlewares

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthenticated",
				"message": "Token Not Found",
			})
			return
		}

		bearer := strings.HasPrefix(token, "Bearer")
		if !bearer {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthenticated",
				"message": "Bearer Not FOund",
			})
			return
		}

		tokenStr := strings.TrimPrefix(token, "Bearer ")

		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthenticated",
				"message": "Token STR",
			})
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)
		if err != nil {
			log.Errorln("ERROR:", err)
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthenticated",
				"error":   err.Error(),
			})
			return
		}

		userId := claims.UserID

		ctx.Set("user_id", userId)
		ctx.Set("phone_number", claims.PhoneNumber)
		ctx.Set("exp_date", claims.ExpiresAt)

		ctx.Next()
	}
}
