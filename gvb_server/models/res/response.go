package res

import (
	"github.com/gin-gonic/gin"
	"gvb_server/utils"
	"net/http"
)

const (
	SUCCESS = 0
	ERROR   = 7
)

// Response RESTful api structure
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// ListResponse 分页数据结构
type ListResponse[T interface{}] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "Success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "Success", c)
}

func OkWithList(list interface{}, count int64, c *gin.Context) {
	OkWithData(ListResponse[interface{}]{Count: count, List: list}, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithError(err error, obj interface{}, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	Result(ERROR, map[string]interface{}{}, "未知错误", c)
}
