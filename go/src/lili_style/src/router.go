package src

import (
	"github.com/gin-gonic/gin"
	"lili_style_test/src/controllers"

)

func GetRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/js", "src/views/js")
	router.Static("/css", "src/views/css")
	router.Static("/img", "src/views/img")
	router.LoadHTMLGlob("src/views/*.html")

	router.GET("/", controllers.GetForm)
	router.GET("/user/login", controllers.GetUserLogin)
	router.POST("/user/login", controllers.PostUserLogin)

	router.GET("/corporate/login", controllers.GetCorporateLogin)
	router.POST("/corporate/login", controllers.PostCorporateLogin)
	router.GET("/corporate/userDetail", controllers.GetCorporateUserDetail)
	router.POST("/corporate/userDetail", controllers.PostCorporateUserDetail)

	return router
}