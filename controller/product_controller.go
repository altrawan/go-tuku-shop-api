package controller

import (
	"net/http"

	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/service"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	List(ctx *gin.Context)
	// Detail(ctx *gin.Context)
	// Store(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
}

type iProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) ProductController {
	return &iProductController{s}
}

// List godoc
// @Summary List products
// @Description get products
// @Tags Product
// @Accept json
// @Produce json
// @Router /product [get]
func (c *iProductController) List(ctx *gin.Context) {
	var products []entity.Product = c.service.List()
	res := helper.BuildSuccessResponse("Success", products)
	ctx.JSON(http.StatusOK, res)
}
