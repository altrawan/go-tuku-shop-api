package controller

import (
	"net/http"

	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/security/token"
	"go-tuku-shop-api/service"

	"github.com/gin-gonic/gin"
)

type StoreController interface {
	Update(ctx *gin.Context)
}

type iStoreController struct {
	service service.StoreService
}

func NewStoreController(s service.StoreService) StoreController {
	return &iStoreController{s}
}

func (c *iStoreController) Update(ctx *gin.Context) {
	var StoreUpdateDTO dto.StoreUpdateDTO
	err := ctx.ShouldBind(&StoreUpdateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id, err := token.ExtractTokenID(ctx)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to extract token", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	StoreUpdateDTO.UserID = id
	data := c.service.Update(StoreUpdateDTO)
	res := helper.BuildSuccessResponse("Success update Store", data)
	ctx.JSON(http.StatusOK, res)
}
