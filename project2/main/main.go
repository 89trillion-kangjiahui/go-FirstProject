package main

import "project2/route"

func main() {
	route.RegisterRoutes().Run(":8080")
}
