package controllers

import (
	"errors"
	"fmt"
	"lili_style_test/sheets"
	"lili_style_test/src/models"
)

func GetAnswers(userdata []string, rowNum int, client *sheets.SheetClient) models.Answers {
	altruism := GetAltruismData(userdata, rowNum, client)
	businessStance := GetBusinessStanceData(userdata, rowNum, client)
	carrierUp := GetCarrierUpData(userdata, rowNum, client)
	commit := GetCommitData(userdata, rowNum, client)
	mentality := GetMentalityData(userdata, rowNum, client)
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

func CheckRateCategory(client *sheets.SheetClient, rate int, row int, category string) error {
	// 行 = row
	// 列 = column
	row += 2
	column := ""
	switch category {
	case "altruism":
		// hv
		column = "hv"
	case "businessStance":
		// hw
		column = "hw"
	case "carrierUp":
		// hx
		column = "hx"
	case "commit":
		// hy
		column = "hy"
	case "mentality":
		// hz
		column = "hz"
	default:
		err := errors.New("categoryに問題があります。")
		return err
	}

	SearchAndSaveRate(client, rate, column, row)
	return nil
}

func SearchAndSaveRate(client *sheets.SheetClient,rate int, column string ,row int){
	err := client.Update(rate, column, row)
	if err != nil {
		fmt.Println(err)
	}
}
