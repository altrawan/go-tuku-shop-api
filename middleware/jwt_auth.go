package middleware

import (
	"net/http"

	"go-tuku-shop-api/helper"
	"go-tuku-shop-api/security/token"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := token.TokenValid(ctx)
		if err != nil {
			err := helper.BuildFailedResponse("Invalid Token", err.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		ctx.Next()
	}
}
