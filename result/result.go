/**
 * @Author: lixiang
 * @Date: 2025/3/12 14:29
 * @Description: result.go
 */

package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功的响应
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	result := &Result{}
	result.Code = ApiCode.Success
	result.Msg = ApiCode.getMsg(ApiCode.Success)
	result.Data = data
	c.JSON(http.StatusOK, result)
}
func Fail(c *gin.Context, message string) {
	result := Result{}
	result.Msg = message
	result.Code = ApiCode.Failed
	result.Data = gin.H{}
	c.JSON(http.StatusOK, result)
}
