package controller

import (
	"go-tuku-shop-api/dto"
	"go-tuku-shop-api/entity"
	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/security/token"
	"go-tuku-shop-api/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Store(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
	UpdatePayment(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type iTransactionController struct {
	service service.TransactionService
}

func NewTransactionController(s service.TransactionService) TransactionController {
	return &iTransactionController{s}
}

// List godoc
// @Summary List transactions
// @Description Get list transactions
// @Tags Transaction
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /transaction [get]
func (c *iTransactionController) List(ctx *gin.Context) {
	var Transactions []entity.Transaction = c.service.List()
	res := helper.BuildSuccessResponse("Success get list Transaction", Transactions)
	ctx.JSON(http.StatusOK, res)
}

// Detail godoc
// @Summary Detail transaction
// @Description Get detail transaction
// @Tags Transaction
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /transaction/{id} [get]
func (c *iTransactionController) Detail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var Transaction entity.Transaction = c.service.FindByPK(id)
	if (Transaction == entity.Transaction{}) {
		res := helper.BuildFailedResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildSuccessResponse("Success get detail profile", Transaction)
	ctx.JSON(http.StatusOK, res)
}

// Store godoc
// @Summary Store transaction
// @Description Post transaction
// @Tags Transaction
// @Param Body body dto.TransactionCreateDTO true "the body to create Transaction"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /transaction [post]
func (c *iTransactionController) Store(ctx *gin.Context) {
	var transactionCreateDTO dto.TransactionCreateDTO
	err := ctx.ShouldBind(&transactionCreateDTO)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		res := helper.BuildFailedResponse("Failed to extract token", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	transactionCreateDTO.UserID = userId
	transactionCreateDTO.Invoice = time.Now().Format("20060102150405")
	transactionCreateDTO.Total = transactionCreateDTO.Qty * transactionCreateDTO.Price
	transactionCreateDTO.Status = "NEW"

	data := c.service.Store(transactionCreateDTO)
	res := helper.BuildSuccessResponse("Success store transaction", data)
	ctx.JSON(http.StatusCreated, res)

}

// Update godoc
// @Summary Update transaction address
// @Description Update transaction address
// @Tags Transaction
// @Param id path string true "id"
// @Param Body body dto.TransactionUpdateAddressDTO true "the body to update Address Transaction"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /transaction/{id}/address [put]
func (c *iTransactionController) UpdateAddress(ctx *gin.Context) {
	var TransactionUpdateAddressDTO dto.TransactionUpdateAddressDTO
	err := ctx.ShouldBind(&TransactionUpdateAddressDTO)
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

	var Transaction entity.Transaction = c.service.FindByPK(id)
	if (Transaction == entity.Transaction{}) {
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

	var transaction entity.Transaction = c.service.FindByUserID(userId)
	if c.service.IsAllowedToEdit(userId, transaction.UserID) {
		TransactionUpdateAddressDTO.ID = Transaction.ID
		data := c.service.UpdateAddress(TransactionUpdateAddressDTO)
		res := helper.BuildSuccessResponse("Success update address", data)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary Update Transaction Payment
// @Description Update transaction payment
// @Tags Transaction
// @Param id path string true "id"
// @Param Body body dto.TransactionUpdatePaymentDTO true "the body to update Payment Transaction"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400,404 {object} helper.Response
// @Router /transaction/{id}/payment [put]
func (c *iTransactionController) UpdatePayment(ctx *gin.Context) {
	var TransactionUpdatePaymentDTO dto.TransactionUpdatePaymentDTO
	err := ctx.ShouldBind(&TransactionUpdatePaymentDTO)
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

	var entityTransaction entity.Transaction = c.service.FindByPK(id)
	if (entityTransaction == entity.Transaction{}) {
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

	var transaction entity.Transaction = c.service.FindByUserID(userId)
	if c.service.IsAllowedToEdit(userId, transaction.UserID) {
		TransactionUpdatePaymentDTO.ID = transaction.ID
		data := c.service.UpdatePayment(TransactionUpdatePaymentDTO)
		res := helper.BuildSuccessResponse("Success update payment", data)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildFailedResponse("You don't have permission", "You are not the owner", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete transaction
// @Description Delete transaction
// @Tags Transaction
// @Param id path string true "id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /transaction/{id} [delete]
func (c *iTransactionController) Delete(ctx *gin.Context) {
	var Transaction entity.Transaction
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildFailedResponse("No param id was not found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	Transaction.ID = id
	c.service.Delete(Transaction)
	res := helper.BuildSuccessResponse("Success delete transaction", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
