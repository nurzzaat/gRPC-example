package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/gRPC-example/auth/models"
	pb "github.com/nurzzaat/gRPC-example/auth/proto"
	"github.com/nurzzaat/gRPC-example/common"
	log "github.com/sirupsen/logrus"
)

func RBACMiddleware(client pb.AuthServiceClient, requiredRole int, requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("userID")

		user, err := client.GetUserRoles(c, &pb.UserID{Id: uint32(userID)})
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorResponse{
				Result: common.ErrorDetail{
					Code:    "99",
					Message: "Failed to get user roles",
				},
			})
			log.Error("Failed to get user roles")
			c.Abort()
			return
		}
		fmt.Println(userID, user)

		log.Infof("userID - %v , roleID - %v", userID, user.Roles)

		hasRole := false
		for _, role := range user.Roles {
			if role == int32(requiredRole) {
				hasRole = true
			}
			if contains(models.Roles[role-1].Permissions, requiredPermission) {
				hasRole = true
			}
		}

		if !hasRole {
			c.JSON(http.StatusBadRequest, common.ErrorResponse{
				Result: common.ErrorDetail{
					Code:    "101",
					Message: "Permission denied",
				},
			})
			log.Errorf("Permission denied: roleID - %+v, userID - %v", user.Roles, userID)
			c.Abort()
			return
		}

		c.Next()
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
