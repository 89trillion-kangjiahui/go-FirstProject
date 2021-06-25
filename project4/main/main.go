package main

import "project3/route"

func main() {
	route.RegisterRoutes().Run(":8000")
}
