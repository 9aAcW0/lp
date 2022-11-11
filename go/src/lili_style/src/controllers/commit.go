package controllers

import (
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)

func GetCommitData(userdata []string) models.Commit {
	answer := userdata[60:70]
	// はいで加点を整形
	yesAdd := answer[3:5]
	yesAdd = append(yesAdd,answer[0])
	// いいえで加点を整形
	noAdd1 := answer[1:3]
	noAdd2 := answer[5:10]
	noAdd := append(noAdd1,noAdd2...)

	yesPointAnswer := utils.ParsedAnswerYesAdd(yesAdd)
	noPointAnswer := utils.ParsedAnswerNoAdd(noAdd)
	ParsedAnswer := append(yesPointAnswer,noPointAnswer...)

	answerInt := utils.ParsedRate(ParsedAnswer, 10)
	commit := GetCommitAnswer(userdata)
	commit.Rate = answerInt
	return commit
}

func GetCommitAnswer(user []string) models.Commit{
	return models.Commit{
		One: 	  user[60],
		Two:	  user[61],
		Three:	  user[62],
		Four:     user[63],
		Five:     user[64],
		Six:      user[65],
		Seven:    user[66],
		Eight:    user[67],
		Nine:     user[68],
		Ten:      user[69],
	}
}