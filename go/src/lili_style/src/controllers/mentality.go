package controllers

import (
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)

func GetMentalityData(userdata []string) models.Mentality {
	answer := userdata[26:48]

	// はいで加点を整形
	var yesAdd []string
	yesAdd = append(yesAdd,answer[1])
	yesAdd = append(yesAdd,answer[4])
	yesAdd = append(yesAdd,answer[9])
	yesAdd = append(yesAdd,answer[11])
	yesAdd = append(yesAdd,answer[15])
	yesAdd = append(yesAdd,answer[19])
	// いいえで加点を整形
	noAdd1 := answer[20:22]
	noAdd2 := answer[16:19]
	noAdd3 := answer[12:15]
	noAdd4 := answer[5:8]
	noAdd4 = append(noAdd4,answer[0])
	noAdd4 = append(noAdd4,answer[10])
	noAdd := append(noAdd1,noAdd2...)
	noAdd = append(noAdd,noAdd3...)
	noAdd = append(noAdd,noAdd4...)

	yesPointAnswer := utils.ParsedAnswerYesAdd(yesAdd)
	noPointAnswer := utils.ParsedAnswerNoAdd(noAdd)
	ParsedAnswer := append(yesPointAnswer,noPointAnswer...)


	answerInt := utils.ParsedRate(ParsedAnswer, 22)
	mentality := GetMentalityAnswer(userdata)
	mentality.Rate = answerInt
	return mentality
}

func GetMentalityAnswer(user []string) models.Mentality {
	return models.Mentality{
		One: 	   user[26],
		Two:	   user[27],
		Three:	   user[28],
		Four:      user[29],
		Five:      user[30],
		Six:       user[31],
		Seven:     user[32],
		Eight:     user[33],
		Nine:      user[34],
		Ten:       user[35],
		Eleven:    user[36],
		Twelve:    user[37],
		Thirteen:  user[38],
		Fourteen:  user[39],
		Fifteen:   user[40],
		Sixteen:   user[41],
		Seventeen: user[42],
		Eighteen:  user[43],
		Nineteen:  user[44],
		Twenty:    user[45],
		TwentyOne: user[46],
		TwentyTwo: user[47],
	}
}