package controllers

import (
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)


func GetAltruismData(userdata []string) models.Altruism {
	answer := userdata[48:60]
	// はいで加点を整形
	yesAdd := answer[10:12]
	yesAdd = append(yesAdd,answer[7])
	yesAdd = append(yesAdd,answer[6])
	yesAdd = append(yesAdd,answer[5])
	yesAdd = append(yesAdd,answer[3])
	yesAdd = append(yesAdd,answer[0])
	// いいえで加点を整形
	noAdd := answer[1:3]
	noAdd = append(noAdd,answer[4])
	noAdd = append(noAdd,answer[8])
	noAdd = append(noAdd,answer[9])

	yesPointAnswer := utils.ParsedAnswerYesAdd(yesAdd)
	noPointAnswer := utils.ParsedAnswerNoAdd(noAdd)
	ParsedAnswer := append(yesPointAnswer,noPointAnswer...)

	answerInt := utils.ParsedRate(ParsedAnswer, 12)
	altruism := GetAltruismAnswer(userdata)
	altruism.Rate = answerInt
	return altruism
}
func GetAltruismAnswer(user []string) models.Altruism {
	return models.Altruism{
		One:	  user[48],
		Two: 	  user[49],
		Three:	  user[50],
		Four:     user[51],
		Five:     user[52],
		Six:      user[53],
		Seven:    user[54],
		Eight:    user[55],
		Nine:     user[56],
		Ten:      user[57],
		Eleven:   user[58],
		Twelve:   user[59],
	}
}