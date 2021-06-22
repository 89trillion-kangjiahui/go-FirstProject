package main

import (
	"project2/controller"

	"github.com/gin-gonic/gin"
)

const url = "/cal/getRet"

func main() {
	r := gin.Default()
	controller.ComputeRet(r, url)
	r.Run(":8000")
}
