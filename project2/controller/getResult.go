package controller

import (
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"

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
	} else if !StringTrue(exp) {
		ret := dto.SetResult(4002, 0, "你输入的表达式算数表达式不正确")
		c.JSON(http.StatusOK, ret)
	} else {
		solution := util.GetResult(util.MixToPost(exp))
		//封装结果
		Ret := dto.SetResult(200, solution, "计算成功")
		c.JSON(http.StatusOK, Ret)
	}
}

// 判断算数表达式是否正确
func StringTrue(exp string) bool {
	//如果算数表达式的第一个位置的元素不是数字或者最后一个位置的元素不是数字返回false
	if !unicode.IsDigit(rune(exp[0])) || !unicode.IsDigit(rune(exp[len(exp)-1])) {
		return false
	}
	flag := false
	for _, char := range exp {
		//如果字符为空，跳过判断
		if char == ' ' {
			continue
		} else if !unicode.IsDigit(rune(char)) && flag == false {
			//如果当前不是数字并且还没有开启运算符检查，则将开启下一个是否为运算符的检测
			flag = true
		} else if !unicode.IsDigit(rune(char)) && flag == true {
			//如果当前不是数字并且开启了运算符检测，则返回false
			return false
		} else if unicode.IsDigit(rune(char)) && flag == true {
			//如果当前是数字并且开启了运算符检测，则关闭运算符检测
			flag = false
		}
	}
	return true
}
