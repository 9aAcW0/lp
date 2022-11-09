package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

)
func GetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func GetUserLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "userLogin.html", nil)
}

func PostUserLogin(c *gin.Context) {
	mail := c.PostForm("mail")
	user, err := GetUser(mail)
	fmt.Println(user)
	if err != nil {
		c.Redirect(301, "/user/login")
		return
	}
	c.HTML(http.StatusOK, "userPage.html", gin.H{"user": user})

}

func GetCorporateLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "corporateLogin.html", nil)
}

func PostCorporateLogin(c *gin.Context) {
	id := c.PostForm("corporateID")
	ps := c.PostForm("password")
	users, err := GetUserByCorporateID(id,ps)
	if err != nil {
		c.Redirect(301, "/corporate/login")


		return
	}

	c.HTML(http.StatusOK, "corporatePage.html", gin.H{"users": users})
}

func GetCorporateUserDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "corporateDetail.html", nil)
}

func PostCorporateUserDetail(c *gin.Context) {
	mail := c.PostForm("mail")
	user, err := GetUserBySlash(mail)
	if err != nil {
		c.Redirect(301, "/")
		return
	}
	c.HTML(http.StatusOK, "corporateDetail.html", gin.H{"user": user})
}