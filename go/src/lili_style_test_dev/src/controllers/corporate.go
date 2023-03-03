package controllers

import (
	"lili_style_test/sheets"
)

func loginCorporate(corporateID string, passcode string) bool {
	formDate, client := sheets.GetCorporateSheet()
	client = nil
	if client == nil {
		for _, d := range formDate {
			if d[1] == corporateID && d[2] == passcode {
				return true
			}
		}
	}
	return false
}
