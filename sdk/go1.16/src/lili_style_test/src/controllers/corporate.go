package controllers

import (
	"fmt"
	"lili_style_test/sheets"
)

func loginCorporate(corporateID string, passcode string) bool {
	formDate := sheets.GetCorporateSheet()
	fmt.Println(corporateID)
	fmt.Println(passcode)
	for _, d := range formDate {
		if d[1] == corporateID && d[2] == passcode {
			return true
		}
	}
	return false
}
