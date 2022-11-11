package controllers

import (
	"lili_style_test/src/models"
)

func GetAnswers(userdata []string) models.Answers {
	altruism := GetAltruismData(userdata)
	businessStance := GetBusinessStanceData(userdata)
	carrierUp := GetCarrierUpData(userdata)
	commit := GetCommitData(userdata)
	mentality := GetMentalityData(userdata)
	personality := GetPersonalityData(userdata)
	return models.Answers{
		Altruism: 		altruism,
		BusinessStance: businessStance,
		CarrierUp: 		carrierUp,
		Commit: 		commit,
		Mentality: 		mentality,
		Personality: 	personality,
	}
}