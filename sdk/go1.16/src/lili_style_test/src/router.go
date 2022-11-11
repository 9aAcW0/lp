package src

import (
	"fmt"
	"lili_style_test/src/controllers"
	"net/http"
	"text/template"
)



func GetForm(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("aaa")
	t, err := template.ParseFiles("src/views/index.gtpl")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}


func GetUserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ujbnfaoij")
	t, err := template.ParseFiles("src/views/userLogin.gtpl")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func PostUserLogin(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("src/views/userPage.gtpl"))
	user, err := controllers.GetUser(r.FormValue("mail"))
	if err != nil {
		panic(err.Error())
	}
	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "userPage.gtpl", user); err != nil {
		fmt.Println(err)
	}
}

func GetCorporateLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/views/corporateLogin.gtpl")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func PostCorporateLogin(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("src/views/corporatePage.gtpl"))
	fmt.Println(r.FormValue("mail"))
	users, err := controllers.GetUserByCorporateID(r.FormValue("corporateID"),r.FormValue("password"))
	if err != nil {
		panic(err.Error())
	}
	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "corporatePage.gtpl", users); err != nil {
		fmt.Println(err)
	}
}

func GetCorporateUserDetail(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/views/corporateDetail.gtpl")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func PostCorporateUserDetail(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("src/views/corporateDetail.gtpl"))
	fmt.Println(r.FormValue("mail"))
	user, err := controllers.GetUser(r.FormValue("mail"))
	if err != nil {
		panic(err.Error())
	}
	if err := tpl.ExecuteTemplate(w, "corporateDetail.gtpl", user); err != nil {
		fmt.Println(err)
	}
}