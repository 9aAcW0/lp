package controllers

import (
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)

func GetCarrierUpData(userdata []string) models.CarrierUp {
	answer := userdata[11:26]
	// はいで加点を整形
	yesAdd := answer[4:13]
	// いいえで加点を整形
	noAdd := answer[0:4]
	noAdd = append(noAdd,answer[13])
	// 年収部分
	annualIncome := 0
	switch answer[14] {
	case"〜300万円":
		annualIncome = 1
	case"300〜500万円":
		annualIncome = 2
	case"500〜800万円":
		annualIncome = 3
	case"800〜1000万円":
		annualIncome = 4
	case"1000万円〜":
		annualIncome = 5
	}

	yesPointAnswer := utils.ParsedAnswerYesAdd(yesAdd)
	noPointAnswer := utils.ParsedAnswerNoAdd(noAdd)
	ParsedAnswer := append(yesPointAnswer,noPointAnswer...)
	ParsedAnswer = append(ParsedAnswer,annualIncome)

	answerInt := utils.ParsedRate(ParsedAnswer, 20)
	CarrierUp := GetCarrierUpAnswer(userdata)
	CarrierUp.Rate = answerInt
	return CarrierUp
}

func GetCarrierUpAnswer(user []string) models.CarrierUp{
	return models.CarrierUp{
		One: 	  user[11],
		Two:	  user[12],
		Three:	  user[12],
		Four:     user[12],
		Five:     user[12],
		Six:      user[12],
		Seven:    user[12],
		Eight:    user[12],
		Nine:     user[12],
		Ten:      user[12],
		Eleven:   user[12],
		Twelve:   user[12],
		Thirteen: user[12],
		Fourteen: user[12],
		Fifteen:  user[25],
	}
}
