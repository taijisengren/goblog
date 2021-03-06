package app

import (
	"github.com/gin-gonic/gin"
	"goblog/pkg/errcode"
	"net/http"
)

// 响应处理
type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx}
}

func (r *Response) ToResponse(data map[string]interface{}) {
	if data == nil {
		data = gin.H{
			"code": 0,
			"msg":  "success",
		}
	} else {
		data["code"] = 0
		data["msg"] = "success"
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code, "msg": err.Msg}
	details := err.Details
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
