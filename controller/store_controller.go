package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/dto"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/helper"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/security/token"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/service"
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
