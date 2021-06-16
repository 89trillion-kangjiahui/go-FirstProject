package main

import (
	"github.com/gin-gonic/gin"
	 . "project2/controller"
)

func main()  {
	r := gin.Default()
	ComputeRet(r)
	r.Run(":8000")
}

