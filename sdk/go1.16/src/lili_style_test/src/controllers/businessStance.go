package controllers

import (
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)

func GetBusinessStanceData(userdata []string) models.BusinessStance {
	answer := userdata[70:83]
	// はいで加点を整形
	yesAdd := answer[0:11]
	yesAdd = append(yesAdd,answer[12])

	// いいえで加点を整
	var noAdd []string
	noAdd = append(noAdd,answer[11])

	yesPointAnswer := utils.ParsedAnswerYesAdd(yesAdd)
	noPointAnswer := utils.ParsedAnswerNoAdd(noAdd)
	ParsedAnswer := append(yesPointAnswer,noPointAnswer...)

	answerInt := utils.ParsedRate(ParsedAnswer, 13)
	BusinessStance := GetBusinessStanceAnswer(userdata)
	BusinessStance.Rate = answerInt
	return BusinessStance
}

func GetBusinessStanceAnswer(user []string) models.BusinessStance {
	return models.BusinessStance{
		One:	  user[70],
		Two: 	  user[71],
		Three: 	  user[72],
		Four:     user[73],
		Five:     user[74],
		Six:      user[75],
		Seven:    user[76],
		Eight:    user[77],
		Nine:     user[78],
		Ten:      user[79],
		Eleven:   user[80],
		Twelve:   user[81],
		Thirteen: user[82],
	}
}