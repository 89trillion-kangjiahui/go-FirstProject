package main

import (
	"github.com/gin-gonic/gin"

	"project2/controller"
)

const url = "/cal/getRet"

func main() {
	r := gin.Default()
	controller.ComputeRet(r, url)
	r.Run(":8000")
}
