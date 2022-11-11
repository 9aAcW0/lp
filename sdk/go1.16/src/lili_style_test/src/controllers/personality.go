package controllers

import (
	"lili_style_test/src/models"
	"lili_style_test/src/utils"
)

func GetPersonalityData(userdata []string) models.Personality {
	answer := GetPersonalityAnswer(userdata)
	answer.RigorRate = GetRigorRate(userdata)
	answer.AcceptabilityRate = GetAcceptabilityRate(userdata)
	answer.LogicalityRate = GetLogicalityRate(userdata)
	answer.FreedomRate = GetFreedomRate(userdata)
	answer.CoordinationRate = GetCoordinationRate(userdata)
	return answer
}

// 厳格度 rigor
func GetRigorRate(userdata []string) int {
	answer := userdata[83:87]
	answerInt := utils.ParsedAnswerPersonality(answer)
	rate := utils.ParsedRate(answerInt, 8)
	return rate
}
// 受容度 acceptability
func GetAcceptabilityRate(userdata []string) int {
	answer := userdata[87:91]
	answerInt := utils.ParsedAnswerPersonality(answer)
	rate := utils.ParsedRate(answerInt, 8)
	return rate
}
// 論理度 logicality
func GetLogicalityRate(userdata []string) int {
	answer := userdata[91:95]
	answerInt := utils.ParsedAnswerPersonality(answer)
	rate := utils.ParsedRate(answerInt, 8)
	return rate
}
// 自由度 degreeOfFreedom
func GetFreedomRate(userdata []string) int {
	answer := userdata[95:99]
	answerInt := utils.ParsedAnswerPersonality(answer)
	rate := utils.ParsedRate(answerInt, 8)
	return rate
}
// 協調度 coordination
func GetCoordinationRate(userdata []string) int {
	answer := userdata[99:]
	answerInt := utils.ParsedAnswerPersonality(answer)
	rate := utils.ParsedRate(answerInt, 8)
	return rate
}

func GetPersonalityAnswer(user []string) models.Personality {
	return models.Personality {
		One:	  		  user[83],
		Two:              user[84],
		Three:            user[85],
		Four:			  user[86],
		Five:			  user[87],
		Six:			  user[88],
		Seven:			  user[89],
		Eight:			  user[90],
		Nine:			  user[91],
		Ten:			  user[92],
		Eleven:			  user[93],
		Twelve:			  user[94],
		Thirteen:		  user[95],
		Fourteen:		  user[96],
		Fifteen:   		  user[97],
		Sixteen:  		  user[98],
		Seventeen:		  user[99],
		Eighteen: 		  user[100],
		Nineteen: 		  user[101],
		Twenty:   		  user[102],
	}


}