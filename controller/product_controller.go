package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/entity"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/helper"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/service"
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
