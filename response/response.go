package response

import (
	"Book-API-Server/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Failed(ctx *gin.Context, err error) {
	if v, ok := err.(*exception.ApiException); ok {
		if v.HttpCode == 0 {
			v.HttpCode = 500
		}
		ctx.JSON(v.HttpCode, v)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "error": err.Error()})
	}
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}
