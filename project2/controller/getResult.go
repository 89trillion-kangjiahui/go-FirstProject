package controller

import (
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"net/http"

	"project2/dto"
	"project2/util"
)

func ComputeRet(c *gin.Context) {
	//go会默认将url中的"+"转化为空格，设置下面的将会获得url中的"+"
	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")
	//从请求中获取表达式
	exp := c.Query("exp")
	if exp == "" {
		ret := dto.SetResult(4001, 0, "你输入的表达式为空")
		c.JSON(http.StatusOK, ret)
	}
	if !StringTrue(exp) {
		ret := dto.SetResult(4002, 0, "你输入的表达式算数表达式不正确")
		c.JSON(http.StatusOK, ret)
	}
	if exp != "" {
		solution := util.GetResult(util.MixToPost(exp))
		//封装结果
		Ret := dto.SetResult(200, solution, "计算成功")
		c.JSON(http.StatusOK, Ret)
	}
}

// 判断算数表达式是否正确
func StringTrue(exp string) bool {
	flag := false
	for _, char := range exp {
		//如果字符为空，跳过判断
		if char == ' ' {
			continue
		} else if !unicode.IsDigit(rune(char)) && flag == false {
			flag = true
		} else if !unicode.IsDigit(rune(char)) && flag == true {
			return false
		} else if unicode.IsDigit(rune(char)) && flag == true {
			flag = false
		}
	}
	if !unicode.IsDigit(rune(exp[len(exp)-1])) {
		return false
	}
	return true
}
