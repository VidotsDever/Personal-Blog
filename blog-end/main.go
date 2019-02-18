package main

import (
	"publisher-end/router"
)

func main() {
	router := router.GetRouter()
	router.Run(":8080")
}
