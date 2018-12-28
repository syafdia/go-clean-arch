package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func Write(ctx *gin.Context, httpStatus int, resp Response) {
	ctx.JSON(httpStatus, resp)
}
