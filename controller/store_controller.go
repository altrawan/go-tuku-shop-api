package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/security/token"
	"go-tuku-shop-api/service"

	"github.com/gin-gonic/gin"
)

type StoreController interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Update(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}

type iStoreController struct {
	service service.StoreService
}

func NewStoreController(s service.StoreService) StoreController {
	return &iStoreController{s}
}

// List godoc
// @Summary List stores
// @Description Get list stores
// @Tags Store
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /store [get]
func (c *iStoreController) List(ctx *gin.Context) {
	var Stores []entity.Store = c.service.List()
	res := helper.BuildSuccessResponse("Success get list Store", Stores)
	ctx.JSON(http.StatusOK, res)
}

// Detail godoc
// @Summary Detail store
// @Description Get detail store
// @Tags Store
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /store/{id} [get]
func (c *iStoreController) Detail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var Store entity.Store = c.service.FindByPK(id)
	if (Store == entity.Store{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildSuccessResponse("Success get detail Store", Store)
	ctx.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary Update store
// @Description Update store
// @Tags Store
// @Param id path string true "id"
// @Param Body body dto.StoreUpdateDTO true "the body to update store"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404,500 {object} helper.Response
// @Router /store/{id} [put]
func (c *iStoreController) Update(ctx *gin.Context) {
	var StoreUpdateDTO dto.StoreUpdateDTO
	err := ctx.ShouldBind(&StoreUpdateDTO)
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

	var store entity.Store = c.service.FindByPK(id)
	if (store == entity.Store{}) {
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

	fmt.Println(userId, store.UserID)

	if c.service.IsAllowedToEdit(userId, store.UserID) {
		StoreUpdateDTO.UserID = userId
		data := c.service.Update(StoreUpdateDTO)
		res := helper.BuildSuccessResponse("Success update Store", data)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

// ChangePassword godoc
// @Summary Change password store
// @Description Change password store
// @Tags Store
// @Param Body body dto.StoreChangePasswordDTO true "the body to change password"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /store/change-password [put]
func (c *iStoreController) ChangePassword(ctx *gin.Context) {
	var StoreChangePasswordDTO dto.StoreChangePasswordDTO
	err := ctx.ShouldBind(&StoreChangePasswordDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if StoreChangePasswordDTO.Password != StoreChangePasswordDTO.ConfirmPassword {
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

	var store entity.Store = c.service.FindByUserID(userId)
	fmt.Println(store.UserID, userId)
	if c.service.IsAllowedToEdit(userId, store.UserID) {
		StoreChangePasswordDTO.UserID = userId
		data := c.service.ChangePassword(StoreChangePasswordDTO)
		res := helper.BuildSuccessResponse("Success change password", data)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
