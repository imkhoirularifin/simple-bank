package controller

import (
	"net/http"
	"simple-bank/db/dto"
	"simple-bank/usecase"
	"simple-bank/utils/common"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginHandler(ctx *gin.Context)
	LogoutHandler(ctx *gin.Context)
	Route()
}

type authController struct {
	authUc usecase.AuthUseCase
	rg     *gin.RouterGroup
	session common.Session
}

func (a *authController) LoginHandler(ctx *gin.Context) {
	var payload dto.LoginParams

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	_, err := a.authUc.Login(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = a.session.StoreSession(payload.Email, ctx)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "login success", nil)
}

func (a *authController) LogoutHandler(ctx *gin.Context) {
	err := a.session.DeleteSession(ctx)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	common.SendSingleResponse(ctx, http.StatusOK, "logout success", nil)
}

func (a *authController) Route() {
	a.rg.POST("/login", a.LoginHandler)
	a.rg.POST("/logout", a.LogoutHandler)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup, session common.Session) AuthController {
	return &authController{
		authUc:  authUc,
		rg:      rg,
		session: session,
	}
}
