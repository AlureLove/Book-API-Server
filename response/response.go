package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Failed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 0, "error": err.Error()})
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}
