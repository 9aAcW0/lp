package main

import "lili_style_test/src"

func main() {
	router := src.GetRouter()
	router.Run(":8080")
}

