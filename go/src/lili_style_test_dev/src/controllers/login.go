package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func GetForm(c *gin.Context) {
	fmt.Println("aaa")
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetUserLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "userLogin.html", nil)
}

func PostUserLogin(c *gin.Context) {
	mail := c.PostForm("mail")
	user, err := GetUser(mail)
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
	var corporate struct{
		ID   string
		Pass string
	}
	corporate.ID = id
	corporate.Pass = ps
	users, err := GetUserByCorporateID(id,ps)
	if err != nil {
		c.Redirect(301, "/corporate/login")
		return
	}

	c.HTML(http.StatusOK, "corporatePage.html", gin.H{"users": users , "corporate": corporate})
}

func GetCorporateUserDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "corporateDetail.html", nil)
}

func PostCorporateUserDetail(c *gin.Context) {
	mail := c.PostForm("mail")
	id := c.PostForm("corporateID")
	ps := c.PostForm("password")
	var corporate struct{
		ID   string
		Pass string
	}
	corporate.ID = id
	corporate.Pass = ps
	user, err := GetUser(mail)
	if err != nil {
		c.Redirect(301, "/")
		return
	}
	c.HTML(http.StatusOK, "corporateDetail.html", gin.H{"user": user, "corporate": corporate})
}