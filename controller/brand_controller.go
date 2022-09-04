package controller

import (
	"net/http"
	"strconv"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/service"

	"github.com/gin-gonic/gin"
)

type BrandController interface {
	List(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type iBrandController struct {
	service service.BrandService
}

func NewBrandController(s service.BrandService) BrandController {
	return &iBrandController{s}
}

func (c *iBrandController) List(ctx *gin.Context) {
	var brands []entity.Brand = c.service.List()
	res := helper.BuildSuccessResponse("Success get list brand", brands)
	ctx.JSON(http.StatusOK, res)
}

func (c *iBrandController) Store(ctx *gin.Context) {
	var brandCreateDTO dto.BrandCreateDTO
	err := ctx.ShouldBind(&brandCreateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		data := c.service.Store(brandCreateDTO)
		res := helper.BuildSuccessResponse("Success store category", data)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (c *iBrandController) Update(ctx *gin.Context) {
	var brandUpdateDTO dto.BrandUpdateDTO
	err := ctx.ShouldBind(&brandUpdateDTO)
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

	var brand entity.Brand = c.service.FindByID(id)
	if (brand == entity.Brand{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	brandUpdateDTO.ID = brand.ID
	data := c.service.Update(brandUpdateDTO)
	res := helper.BuildSuccessResponse("Success update category", data)
	ctx.JSON(http.StatusOK, res)
}

func (c *iBrandController) Delete(ctx *gin.Context) {
	var brand entity.Brand
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	brand.ID = id
	c.service.Delete(brand)
	res := helper.BuildSuccessResponse("Success delete category", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
