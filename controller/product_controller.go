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

type ProductController interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type iProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) ProductController {
	return &iProductController{s}
}

// List godoc
// @Summary List products
// @Description Get list products
// @Tags Brand
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /product [get]
func (c *iProductController) List(ctx *gin.Context) {
	var Products []entity.Product = c.service.List()
	res := helper.BuildSuccessResponse("Success get list Product", Products)
	ctx.JSON(http.StatusOK, res)
}

// Detail godoc
// @Summary Detail product
// @Description Get detail product
// @Tags Product
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /product/{id} [get]
func (c *iProductController) Detail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var product entity.Product = c.service.FindByPK(id)
	if (product == entity.Product{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildSuccessResponse("Success get detail profile", product)
	ctx.JSON(http.StatusOK, res)
}

// Store godoc
// @Summary Store product
// @Description Post product
// @Tags Product
// @Param Body body dto.ProductCreateDTO true "the body to create product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /product [post]
func (c *iProductController) Store(ctx *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	err := ctx.ShouldBind(&productCreateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		data := c.service.Store(productCreateDTO)
		res := helper.BuildSuccessResponse("Success store product", data)
		ctx.JSON(http.StatusCreated, res)
	}
}

// Update godoc
// @Summary Update product
// @Description Update product
// @Tags Product
// @Param id path string true "id"
// @Param Body body dto.ProductUpdateDTO true "the body to update product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /product/{id} [put]
func (c *iProductController) Update(ctx *gin.Context) {
	var ProductUpdateDTO dto.ProductUpdateDTO
	err := ctx.ShouldBindJSON(&ProductUpdateDTO)
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

	var Product entity.Product = c.service.FindByPK(id)
	if (Product == entity.Product{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ProductUpdateDTO.ID = Product.ID
	data := c.service.Update(ProductUpdateDTO)
	res := helper.BuildSuccessResponse("Success update product", data)
	ctx.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete product
// @Description Delete product
// @Tags Product
// @Param id path string true "id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /product/{id} [delete]
func (c *iProductController) Delete(ctx *gin.Context) {
	var Product entity.Product
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	Product.ID = id
	c.service.Delete(Product)
	res := helper.BuildSuccessResponse("Success delete product", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
