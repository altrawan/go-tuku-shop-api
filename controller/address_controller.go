package controller

import (
	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/security/token"
	"go-tuku-shop-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressController interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type iAddressController struct {
	service service.AddressService
}

func NewAddressController(s service.AddressService) AddressController {
	return &iAddressController{s}
}

// List godoc
// @Summary List Addresss
// @Description Get list Addresss
// @Tags Address
// @Param Body body dto.AddressCreateDTO true "the body to create Address"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /Address [get]
func (c *iAddressController) List(ctx *gin.Context) {
	var Addresss []entity.Address = c.service.List()
	res := helper.BuildSuccessResponse("Success get list Address", Addresss)
	ctx.JSON(http.StatusOK, res)
}

// Detail godoc
// @Summary Detail Address
// @Description Get detail Address
// @Tags Address
// @Param Body body dto.AddressCreateDTO true "the body to create Address"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /Address/{id} [get]
func (c *iAddressController) Detail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var Address entity.Address = c.service.FindByPK(id)
	if (Address == entity.Address{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildSuccessResponse("Success get detail profile", Address)
	ctx.JSON(http.StatusOK, res)
}

// Store godoc
// @Summary Store Address
// @Description Post Address
// @Tags Address
// @Param Body body dto.AddressCreateDTO true "the body to create Address"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 201 {object} helper.Response
// @Failure 400,500 {object} helper.Response
// @Router /Address [post]
func (c *iAddressController) Store(ctx *gin.Context) {
	var AddressCreateDTO dto.AddressCreateDTO
	err := ctx.ShouldBind(&AddressCreateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		userId, err := token.ExtractTokenID(ctx)
		if err != nil {
			res := helper.BuildFailedResponse("Failed to extract token", err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}

		AddressCreateDTO.UserID = userId
		data := c.service.Store(AddressCreateDTO)
		res := helper.BuildSuccessResponse("Success store address", data)
		ctx.JSON(http.StatusCreated, res)
	}
}

// Update godoc
// @Summary Update Address
// @Description Update Address
// @Tags Address
// @Param id path string true "id"
// @Param Body body dto.AddressUpdateDTO true "the body to update Address"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404,500 {object} helper.Response
// @Router /Address/{id} [put]
func (c *iAddressController) Update(ctx *gin.Context) {
	var AddressUpdateDTO dto.AddressUpdateDTO
	err := ctx.ShouldBind(&AddressUpdateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var Address entity.Address = c.service.FindByPK(id)
	if (Address == entity.Address{}) {
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

	var address entity.Address = c.service.FindByUserID(userId)
	if c.service.IsAllowedToEdit(userId, address.UserID) {
		AddressUpdateDTO.ID = id
		AddressUpdateDTO.UserID = userId
		data := c.service.Update(AddressUpdateDTO)
		res := helper.BuildSuccessResponse("Success update address", data)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete Address
// @Description Delete Address
// @Tags Address
// @Param id path string true "id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,500 {object} helper.Response
// @Router /Address/{id} [delete]
func (c *iAddressController) Delete(ctx *gin.Context) {
	var Address entity.Address
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to extract token", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	var address entity.Address = c.service.FindByUserID(userId)
	if c.service.IsAllowedToEdit(userId, address.UserID) {
		Address.ID = id
		c.service.Delete(Address)
		res := helper.BuildSuccessResponse("Success delete address", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
