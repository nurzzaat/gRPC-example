package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/nurzzaat/gRPC-example/auth/proto"
	"github.com/nurzzaat/gRPC-example/common"
	log "github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary	SignIn
// @Tags		auth
// @Accept		json
// @Produce	json
// @Param		input	body    LoginRequest	true	"login"
// @Success	200		{object}	common.SuccessResponse
// @Failure	default	{object}	common.ErrorResponse
// @Router		/auth/sign-in [post]
func (lc *AuthController) Signin(c *gin.Context) {
	logFields := log.Fields{
		"requestType": "POST",
		"endpoint":    "/auth/admin/sign-in",
	}

	var loginRequest LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{
			Result: common.ErrorDetail{
				Code:    "100",
				Message: "General data binding error",
			},
		})
		log.WithFields(logFields).Error("General data binding error:", err.Error())
		return
	}
	log.WithFields(logFields).Infof("request from user: %+v", loginRequest)

	response, err := lc.Client.SignIn(c, &pb.SignInRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorDetail{
			Code:    err.Error(),
			Message: "Error from gRPC server",
		})
		return
	}
	log.WithFields(logFields).Infof("response from server: %+v", common.SuccessResponse{Result: response})
	c.JSON(http.StatusOK, common.SuccessResponse{Result: response})
}
