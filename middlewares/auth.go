package middlewares

import (
	"strconv"
	"time"

	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Token Not Found",
			})
			return
		}

		bearer := strings.HasPrefix(token, "Bearer")

		if !bearer {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Bearer Not FOund",
			})
			return
		}
		tokenStr := strings.Split(token, "Bearer ")[1]

		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Token STR",
			})
			return
		}

		claims, err := helpers.VerifyToken(tokenStr)
		if err != nil {
			log.Errorln("ERROR:", err)
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		var data = claims.(jwt.MapClaims)
		userId := data["id"].(float64)
		strUserId := strconv.FormatFloat(userId, 'f', -1, 64)

		ctx.Set("user_id", strUserId)
		ctx.Set("name", data["name"])
		ctx.Set("email", data["email"])
		ctx.Set("exp", data["exp"])
		ctx.Set("exp_date", data["exp_date"])

		if data["exp_date"] == nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
			return
		}

		timeNow := time.Now()
		expiredTime := data["exp_date"].(string)

		parsed, _ := time.Parse(time.RFC3339, expiredTime)

		if err != nil {
			log.Errorln("ERROR:", err)
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		if timeNow.After(parsed) {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error":         "loggedOut",
				"message":       "Token has expired, please login again",
				"is_logged_out": true,
			})
			return
		}

		ctx.Next()
	}
}
