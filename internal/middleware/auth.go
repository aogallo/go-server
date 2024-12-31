package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aogallo/go-server/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Missing Authentication")
			ctx.Abort()
			return
		}

		tokenWithBear := strings.Split(tokenString, " ")

		if len(tokenWithBear) != 2 || tokenWithBear[0] != "Bearer" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid authentication token")
			ctx.Abort()
			return
		}

		tokenString = tokenWithBear[1]

		user, error := utils.VerifyToken(tokenString)

		if error != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, fmt.Sprintf("Invalid authentication token. %s", error.Error()))
			ctx.Abort()
			return
		}

		ctx.Set("user_id", user.ID)
		ctx.Next()

	}
}
