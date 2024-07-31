package middlewares

import (
	"fmt"

	"github.com/DeniesKresna/sined/service/extensions/terror"
	"github.com/gin-gonic/gin"
)

func CheckPaymentGateway() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("waduh")
		var terr terror.ErrInterface
		// Get the authorization header from the request
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		if authHeader != "paymentgatewaytoken" {
			terr = terror.ErrInvalidRule("this operation is not allowed")
			responseJson(ctx, terr)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
