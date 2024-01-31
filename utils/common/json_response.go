package common

import (
	modelutil "simple-bank/utils/model_util"

	"github.com/gin-gonic/gin"
)

func SendSingleResponse(ctx *gin.Context, code int, description string, data any) {
	ctx.JSON(code, modelutil.SingleResponse{
		Status: modelutil.Status{
			Code:        code,
			Description: description,
		},
		Data: data,
	})
}

func SendPagedResponse(ctx *gin.Context, code int, description string, data any, paging any) {
	ctx.JSON(code, modelutil.PagedResponse{
		Status: modelutil.Status{
			Code:        code,
			Description: description,
		},
		Data:   data,
		Paging: paging,
	})
}
