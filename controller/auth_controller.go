package controller

import (
	"net/http"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/security/token"
	"go-tuku-shop-api/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	RegisterSeller(ctx *gin.Context)
	RegisterBuyer(ctx *gin.Context)
	// Activation(ctx *gin.Context)
	// ForgotPassword(ctx *gin.Context)
	// ResetPassword(ctx *gin.Context)
}

type iAuthController struct {
	service service.AuthService
}

func NewAuthController(s service.AuthService) AuthController {
	return &iAuthController{s}
}

// Login godoc
// @Summary Login
// @Description Logging in to get jwt to access admin, seller or buyer api by roles
// @Tags Auth
// @Param Body body dto.LoginDTO true "the body to login a user"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,403 {object} helper.Response
// @Router /auth/login [post]
func (c *iAuthController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO

	if err := ctx.ShouldBind(&loginDTO); err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.service.Login(loginDTO.Email, loginDTO.Password)
	if u, ok := result.(entity.User); ok {
		generatedToken, err := token.GenerateToken(u.ID)
		if err != nil {
			err := helper.BuildFailedResponse("Failed Generate Token", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		user := map[string]string{
			"email": u.Email,
			"token": generatedToken,
		}

		res := helper.BuildSuccessResponse("Login Success", user)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
}

// RegisterSeller godoc
// @Summary Register seller
// @Description Registering a seller
// @Tags Auth
// @Param Body body dto.RegisterSellerDTO true "the body to register a seller"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /auth/register-seller [post]
func (c *iAuthController) RegisterSeller(ctx *gin.Context) {
	var registerSellerDTO dto.RegisterSellerDTO

	if err := ctx.ShouldBind(&registerSellerDTO); err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if !c.service.IsDuplicateEmail(registerSellerDTO.Email) {
		res := helper.BuildFailedResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	} else {
		createdUser := c.service.RegisterSeller(registerSellerDTO)
		res := helper.BuildSuccessResponse("Register Success", createdUser)
		ctx.JSON(http.StatusOK, res)
	}
}

// RegisterBuyer godoc
// @Summary Register buyer
// @Description Registering a buyer
// @Tags Auth
// @Param Body body dto.RegisterBuyerDTO true "the body to register a buyer"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /auth/register-buyer [post]
func (c *iAuthController) RegisterBuyer(ctx *gin.Context) {
	var registerBuyerDTO dto.RegisterBuyerDTO

	if err := ctx.ShouldBind(&registerBuyerDTO); err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if !c.service.IsDuplicateEmail(registerBuyerDTO.Email) {
		res := helper.BuildFailedResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	} else {
		createdUser := c.service.RegisterBuyer(registerBuyerDTO)
		res := helper.BuildSuccessResponse("Register Success", createdUser)
		ctx.JSON(http.StatusOK, res)
	}
}
