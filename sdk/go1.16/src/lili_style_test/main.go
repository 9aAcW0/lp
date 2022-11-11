package main


import (
	"fmt"
	"lili_style_test/src"

	//"lili_style_test/src"
	//"lili_style_test/src/controllers"
	"net/http"
)

func main() {
	fmt.Println("aaa")
	http.HandleFunc("/", src.GetForm)    // `/items`の処理（）
	http.HandleFunc("/user/login", userLoginHandler)
	http.HandleFunc("/corporate/login", corporateLoginHandler)
	http.HandleFunc("/corporate/userDetail", corporateDetailHandler)
	http.ListenAndServe(":8000", nil)
}


func userLoginHandler (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		src.GetUserLogin(w, r) // 全てのitemの取得
	case "POST":
		src.PostUserLogin(w, r) // 新しいitemの追加
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func corporateLoginHandler (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		src.GetCorporateLogin(w, r) // 全てのitemの取得
	case "POST":
		src.PostCorporateLogin(w, r) // 新しいitemの追加
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func corporateDetailHandler (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		src.GetCorporateUserDetail(w, r) // 全てのitemの取得
	case "POST":
		src.PostCorporateUserDetail(w, r) // 新しいitemの追加
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}