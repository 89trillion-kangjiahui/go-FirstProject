package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "project2/dto"
	. "project2/util"
	"strings"
)

func ComputeRet(r *gin.Engine)  {
	r.GET("/cal/getRet", func(c *gin.Context) {
		//go会默认将url中的"+"转化为空格，设置下面的将会获得url中的"+"
		c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")
		//从请求中获取表达式
		exp := c.Query("exp")
		solution := GetResult(MixToPost(exp))
		//封装结果
		ret := *GetCalRet(exp, solution)
		c.JSON(http.StatusOK, ret)
	})
}
