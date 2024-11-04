package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/gRPC-example/common"
	"github.com/nurzzaat/gRPC-example/auth/tokenutil"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func JWTAuth(secret string, redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenutil.ValidateJWT(c, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.ErrorResponse{
				Result: common.ErrorDetail{
					Code:    "102",
					Message: "Authorization error",
					Metadata: common.Properties{
						Properties1: err.Error(),
					},
				},
			})
			log.Error("Authorization error:", err.Error())
			c.Abort()
			return
		}
		err = tokenutil.ValidateUserJWT(c, secret, redisClient)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.ErrorResponse{
				Result: common.ErrorDetail{
					Code:    "103",
					Message: "User is required",
					Metadata: common.Properties{
						Properties1: err.Error(),
					},
				},
			})
			log.Error("User is required:", err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
