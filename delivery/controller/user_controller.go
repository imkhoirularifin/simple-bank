package controller

import (
	"net/http"
	db "simple-bank/db/sqlc"
	"simple-bank/usecase"
	"simple-bank/utils/common"

	"github.com/emicklei/pgtalk/convert"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateHandler(ctx *gin.Context)
	GetHandler(ctx *gin.Context)
	ListHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
	Route()
}

type userController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) UserController {
	return &userController{
		uc: uc,
		rg: rg,
	}
}

func (u *userController) CreateHandler(ctx *gin.Context) {
	var payload db.CreateUserParams

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := u.uc.CreateUser(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "user created", response)
}

func (u *userController) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := u.uc.GetUser(id)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "user found", response)
}

func (u *userController) ListHandler(ctx *gin.Context) {
	response, err := u.uc.ListUsers()
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "users found", response)
}

func (u *userController) UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload db.UpdateUserParams

	payload.ID = convert.StringToUUID(id)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := u.uc.UpdateUser(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "user updated", response)
}

func (u *userController) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := u.uc.DeleteUser(id); err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "user deleted", nil)
}

func (u *userController) Route() {
	u.rg.POST("/users", u.CreateHandler)
	u.rg.GET("/users/:id", u.GetHandler)
	u.rg.GET("/users", u.ListHandler)
	u.rg.PUT("/users/:id", u.UpdateHandler)
	u.rg.DELETE("/users/:id", u.DeleteHandler)
}
