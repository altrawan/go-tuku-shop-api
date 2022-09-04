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

type CategoryController interface {
	List(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type iCategoryController struct {
	service service.CategoryService
}

func NewCategoryController(s service.CategoryService) CategoryController {
	return &iCategoryController{s}
}

func (c *iCategoryController) List(ctx *gin.Context) {
	var Categorys []entity.Category = c.service.List()
	res := helper.BuildSuccessResponse("Success get list Category", Categorys)
	ctx.JSON(http.StatusOK, res)
}

func (c *iCategoryController) Store(ctx *gin.Context) {
	var CategoryCreateDTO dto.CategoryCreateDTO
	err := ctx.ShouldBind(&CategoryCreateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		data := c.service.Store(CategoryCreateDTO)
		res := helper.BuildSuccessResponse("Success store category", data)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (c *iCategoryController) Update(ctx *gin.Context) {
	var CategoryUpdateDTO dto.CategoryUpdateDTO
	err := ctx.ShouldBind(&CategoryUpdateDTO)
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

	var Category entity.Category = c.service.FindByID(id)
	if (Category == entity.Category{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	CategoryUpdateDTO.ID = id
	data := c.service.Update(CategoryUpdateDTO)
	res := helper.BuildSuccessResponse("Success update category", data)
	ctx.JSON(http.StatusOK, res)
}

func (c *iCategoryController) Delete(ctx *gin.Context) {
	var Category entity.Category
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	Category.ID = id
	c.service.Delete(Category)
	res := helper.BuildSuccessResponse("Success delete category", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
