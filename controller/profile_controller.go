package controller

import (
	"net/http"
	"strconv"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/security/token"
	"go-tuku-shop-api/service"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Update(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}

type iProfileController struct {
	service service.ProfileService
}

func NewProfileController(s service.ProfileService) ProfileController {
	return &iProfileController{s}
}

// List godoc
// @Summary List profiles
// @Description Get list profiles
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /profile [get]
func (c *iProfileController) List(ctx *gin.Context) {
	var profiles []entity.Profile = c.service.List()
	res := helper.BuildSuccessResponse("Success get list profile", profiles)
	ctx.JSON(http.StatusOK, res)
}

// Detail godoc
// @Summary Detail profile
// @Description Get detail profile
// @Tags Profile
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /profile/{id} [get]
func (c *iProfileController) Detail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var profile entity.Profile = c.service.FindByPK(id)
	if (profile == entity.Profile{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildSuccessResponse("Success get detail profile", profile)
	ctx.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary Update profile
// @Description Update profile
// @Tags Profile
// @Param id path string true "id"
// @Param Body body dto.ProfileUpdateDTO true "the body to update profile"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404,500 {object} helper.Response
// @Router /profile/{id} [put]
func (c *iProfileController) Update(ctx *gin.Context) {
	var profileUpdateDTO dto.ProfileUpdateDTO
	err := ctx.ShouldBind(&profileUpdateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 60)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var profile entity.Profile = c.service.FindByPK(id)
	if (profile == entity.Profile{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to extract token", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	if c.service.IsAllowedToEdit(userId, profile.UserID) {
		profileUpdateDTO.UserID = userId
		data := c.service.Update(profileUpdateDTO)
		res := helper.BuildSuccessResponse("Success update profile", data)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

// ChangePassword godoc
// @Summary Change password profile
// @Description Change password profile
// @Tags Profile
// @Param Body body dto.ProfileChangePasswordDTO true "the body to change password"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404,500 {object} helper.Response
// @Router /profile/change-password [put]
func (c *iProfileController) ChangePassword(ctx *gin.Context) {
	var profileChangePasswordDTO dto.ProfileChangePasswordDTO
	err := ctx.ShouldBind(&profileChangePasswordDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if profileChangePasswordDTO.Password != profileChangePasswordDTO.ConfirmPassword {
		res := helper.BuildFailedResponse("Confirm password does not match password", "Bad Request", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to extract token", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	var profile entity.Profile = c.service.FindByUserID(userId)
	if c.service.IsAllowedToEdit(userId, profile.UserID) {
		profileChangePasswordDTO.UserID = userId
		data := c.service.ChangePassword(profileChangePasswordDTO)
		res := helper.BuildSuccessResponse("Success change password", data)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
