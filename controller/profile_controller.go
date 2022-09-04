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

func (c *iProfileController) List(ctx *gin.Context) {
	var profiles []entity.Profile = c.service.List()
	res := helper.BuildSuccessResponse("Success get list profile", profiles)
	ctx.JSON(http.StatusOK, res)
}

func (c *iProfileController) Detail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var profile entity.Profile = c.service.FindByID(id)
	if (profile == entity.Profile{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildSuccessResponse("Success get detail profile", profile)
	ctx.JSON(http.StatusOK, res)
}

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

	var profile entity.Profile = c.service.FindByID(id)
	if (profile == entity.Profile{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to extract token", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
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

func (c *iProfileController) ChangePassword(ctx *gin.Context) {

}
