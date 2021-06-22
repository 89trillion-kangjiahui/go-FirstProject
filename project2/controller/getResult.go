package controller

import (
	"strings"

	"project2/dto"
	"project2/util"

	"github.com/gin-gonic/gin"
	"net/http"
)

func ComputeRet(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		//go会默认将url中的"+"转化为空格，设置下面的将会获得url中的"+"
		c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")
		//从请求中获取表达式
		exp := c.Query("exp")
		if exp == "" {
			ret := dto.SetResult(4001, 0, "你输入的表达式为空")
			c.JSON(http.StatusOK, ret)
		}
		if exp != "" {
			solution := util.GetResult(util.MixToPost(exp))
			//封装结果
			Ret := dto.SetResult(200, solution, "计算成功")
			c.JSON(http.StatusOK, Ret)
		}
	})
}
