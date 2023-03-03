package main

import (
	"lili_style_test/src"
)

func main() {
	router := src.GetRouter()
	//port := os.Getenv("PORT")
	//router.Run(":" + port)
	router.Run(":8000")
}

