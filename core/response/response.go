package response

import (
	"github.com/gin-gonic/gin"
	"github.com/syafdia/go-clean-arch/core/logger"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"error,omitempty"`
}

func WriteData(ctx *gin.Context, httpStatus int, data interface{}) {
	write(ctx, httpStatus, Response{Data: data})
}

func WriteError(ctx *gin.Context, httpStatus int, err error) {
	logger.Error(err)
	write(ctx, httpStatus, Response{Error: err})
}

func write(ctx *gin.Context, httpStatus int, resp Response) {
	ctx.JSON(httpStatus, resp)
}
