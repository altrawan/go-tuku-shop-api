package controller

import (
	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type iCartController struct {
	service service.CartService
}

func NewCartController(s service.CartService) CartController {
	return &iCartController{s}
}

// List godoc
// @Summary List carts
// @Description Get list carts
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /cart [get]
func (c *iCartController) List(ctx *gin.Context) {
	var Carts []entity.Cart = c.service.List()
	res := helper.BuildSuccessResponse("Success get list Cart", Carts)
	ctx.JSON(http.StatusOK, res)
}

// Detail godoc
// @Summary Detail cart
// @Description Get detail cart
// @Tags Cart
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /cart/{id} [get]
func (c *iCartController) Detail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var cart entity.Cart = c.service.FindByPK(id)
	if (cart == entity.Cart{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildSuccessResponse("Success get detail profile", cart)
	ctx.JSON(http.StatusOK, res)
}

// Store godoc
// @Summary Store cart
// @Description Post cart
// @Tags Cart
// @Param Body body dto.CartCreateDTO true "the body to create cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /cart [post]
func (c *iCartController) Store(ctx *gin.Context) {
	var CartCreateDTO dto.CartCreateDTO
	err := ctx.ShouldBind(&CartCreateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		data := c.service.Store(CartCreateDTO)
		res := helper.BuildSuccessResponse("Success store category", data)
		ctx.JSON(http.StatusCreated, res)
	}
}

// Update godoc
// @Summary Update cart
// @Description Update cart
// @Tags Cart
// @Param id path string true "id"
// @Param Body body dto.CartUpdateDTO true "the body to update cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /cart/{id} [put]
func (c *iCartController) Update(ctx *gin.Context) {
	var CartUpdateDTO dto.CartUpdateDTO
	err := ctx.ShouldBind(&CartUpdateDTO)
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

	var Cart entity.Cart = c.service.FindByPK(id)
	if (Cart == entity.Cart{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	CartUpdateDTO.ID = Cart.ID
	data := c.service.Update(CartUpdateDTO)
	res := helper.BuildSuccessResponse("Success update category", data)
	ctx.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete cart
// @Description Delete cart
// @Tags Cart
// @Param id path string true "id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /cart/{id} [delete]
func (c *iCartController) Delete(ctx *gin.Context) {
	var Cart entity.Cart
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	Cart.ID = id
	c.service.Delete(Cart)
	res := helper.BuildSuccessResponse("Success delete category", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
