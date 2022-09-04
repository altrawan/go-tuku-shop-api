package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/dto"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/helper"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/security/token"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/service"
)

type ProfileController interface {
	Update(ctx *gin.Context)
}

type iProfileController struct {
	service service.ProfileService
}

func NewProfileController(s service.ProfileService) ProfileController {
	return &iProfileController{s}
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
