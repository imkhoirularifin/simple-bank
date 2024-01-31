package controller

import (
	"net/http"
	db "simple-bank/db/sqlc"
	"simple-bank/delivery/middleware"
	"simple-bank/usecase"
	"simple-bank/utils/common"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	CreateHandler(ctx *gin.Context)
	GetHandler(ctx *gin.Context)
	Route()
}

type transactionController struct {
	tc             usecase.TransactionUseCase
	userUc         usecase.UserUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	session        common.Session
}

func (t *transactionController) CreateHandler(ctx *gin.Context) {
	var payload db.CreateTransactionParams

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	email, err := t.session.ReadSession(ctx)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// get user id from email
	user, err := t.userUc.GetUserWithEmail(email)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	payload.FromUserID = user.ID

	response, err := t.tc.CreateTransaction(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "transaction created", response)
}

func (t *transactionController) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := t.tc.GetTransaction(id)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "transaction found", response)
}

func (t *transactionController) Route() {
	t.rg.POST("/transactions", t.authMiddleware.IsAuthenticated(), t.CreateHandler)
	t.rg.GET("/transactions/:id", t.GetHandler)
}

func NewTransactionController(tc usecase.TransactionUseCase, userUc usecase.UserUseCase, rg *gin.RouterGroup, authMidleware middleware.AuthMiddleware, session common.Session) TransactionController {
	return &transactionController{
		tc:             tc,
		rg:             rg,
		authMiddleware: authMidleware,
		userUc:         userUc,
		session:        session,
	}
}
