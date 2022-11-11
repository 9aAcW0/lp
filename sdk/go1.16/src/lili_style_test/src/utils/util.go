package utils

import (
	"lili_style_test/sheets"
)

func SearchUserDataByMail(mail string) []string {
	formData := sheets.GetAll()
	for _, d := range formData {
		if d[1] == mail {
			user := convertInterfaceToString(d)
			return user
		}
	}
	return nil
}

func SearchUserDataBySlash(mail string) []string {
	formData := sheets.GetAll()
	for _, d := range formData {
		converted := convertInterfaceToString(d)
		if converted[1] + "/" == mail {
			return converted
		}
	}
	return nil
}

func convertInterfaceToString(iData []interface{}) []string{
	var converted []string
	for _, i := range iData {
		converted = append(converted, i.(string))
	}
	return converted
}


func ParsedRate(answers []int, questionCount int) int {
	point := 0
	for _, a := range answers{
		point = point + a
	}
	var coefficient float64
	coefficient = float64(questionCount) / float64(10) //問題数が10の何倍かを求める
	coefficient = 1 / coefficient     //求めた結果で1を割り、係数にする
	result := float64(point) * coefficient	  //係数を合計にかけて最大値を10に近くする。
	return int(result)
}

func ParsedAnswerYesAdd(answers []string) []int{
	var intAnswers []int
	for _, a := range answers {
		switch a {
		case "はい":
			intAnswers = append(intAnswers, 1)
		case "いいえ":
			intAnswers = append(intAnswers, 0)
		}
	}
	return intAnswers
}
func ParsedAnswerNoAdd(answers []string) []int {
	var intAnswers []int
	for _, a := range answers {
		switch a {
		case "はい":
			intAnswers = append(intAnswers, 0)
		case "いいえ":
			intAnswers = append(intAnswers, 1)
		}
	}
	return intAnswers
}

func ParsedAnswerPersonality(answers []string) []int {
	var intAnswers []int
	for _, a := range answers {
		switch a{
		case "はい":
			intAnswers = append(intAnswers, 2)
		case "どちらでもない":
			intAnswers = append(intAnswers, 1)
		case "いいえ":
			intAnswers = append(intAnswers, 0)
		}
	}
	return intAnswers
}

func SearchUserDataByCorporateID(corporateID string) [][]string {
	formData := sheets.GetAll()
	var users [][]string
	for _, d := range formData {
		if d[2] == corporateID {
			user := convertInterfaceToString(d)
			users = append(users, user)
		}
	}
	return users
}

