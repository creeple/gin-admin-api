/**
 * @Author: lixiang
 * @Date: 2025/3/12 14:55
 * @Description: testApi.go
 */

package api

import (
	"gin-admin-api/result"
	"github.com/gin-gonic/gin"
)

func TestSuccess(c *gin.Context) {
	result.Success(c, "success")
}
func TestFail(c *gin.Context) {
	result.Fail(c, "fail")
}
