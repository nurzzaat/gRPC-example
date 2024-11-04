package auth

import (
	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
	pb "github.com/nurzzaat/gRPC-example/auth/proto"
	"github.com/nurzzaat/gRPC-example/common"
)

type Signup struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	verifier = emailverifier.NewVerifier()
)

// @Summary		SignUp
// @Tags			auth
// @Description	signup
// @ID				signup
// @Accept			json
// @Produce		json
// @Param			input	body		Signup	true	"signup"
// @Success		200		{object}	common.SuccessResponse
// @Failure		500		{object}	common.ErrorDetail
// @Failure		default	{object}	common.ErrorDetail
// @Router	 /auth/sign-up [post]
func (sc *AuthController) Signup(c *gin.Context) {
	var request Signup

	verifier = verifier.EnableSMTPCheck()
	verifier = verifier.EnableDomainSuggest()

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorDetail{

			Code:    "NOT_CORRECT_REQUST_FOR_SIGNUP",
			Message: err.Error(),
		})
		return
	}

	response, err := sc.Client.SignUp(c, &pb.SignUpRequest{Email: request.Email, Password: request.Password})
	if err != nil{
		c.JSON(http.StatusBadRequest, common.ErrorDetail{
			Code:    err.Error(),
			Message: "Error from gRPC server",
		})
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse{Result: response})
}
