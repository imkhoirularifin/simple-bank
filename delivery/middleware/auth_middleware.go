package middleware

import (
	"net/http"
	"simple-bank/utils/common"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	IsAuthenticated() gin.HandlerFunc
}

type authMiddleware struct {
	session common.Session
}

func (a *authMiddleware) IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := a.session.ReadSession(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Next()
	}
}

func NewAuthMiddleware(session common.Session) AuthMiddleware {
	return &authMiddleware{session: session}
}
