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

// List godoc
// @Summary List categories
// @Description Get list categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /category [get]
func (c *iCategoryController) List(ctx *gin.Context) {
	var Categorys []entity.Category = c.service.List()
	res := helper.BuildSuccessResponse("Success get list Category", Categorys)
	ctx.JSON(http.StatusOK, res)
}

// Store godoc
// @Summary Store category
// @Description Post category
// @Tags Category
// @Param Body body dto.CategoryCreateDTO true "the body to create category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /category [post]
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

// Update godoc
// @Summary Update category
// @Description Update category
// @Tags Category
// @Param id path string true "id"
// @Param Body body dto.CategoryUpdateDTO true "the body to update category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /category/{id} [put]
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

	var Category entity.Category = c.service.FindByPK(id)
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

// Delete godoc
// @Summary Delete category
// @Description Delete category
// @Tags Category
// @Param id path string true "id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /category/{id} [delete]
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
