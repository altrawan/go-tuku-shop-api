package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/helper"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/security/token"
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
