package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/dto"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/helper"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/security/token"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/service"
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
