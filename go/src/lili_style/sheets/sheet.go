package sheets

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"

	"io/ioutil"
)

type SheetClient struct {
	srv           *sheets.Service
	spreadsheetID string
}


func NewSheetClient(ctx context.Context, spreadsheetID string) (*SheetClient, error) {
	b, err := ioutil.ReadFile("secret.json")
	if err != nil {
		return nil, err
	}
	// read & write permission
	jwt, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}
	srv, err := sheets.New(jwt.Client(ctx))
	if err != nil {
		return nil, err
	}
	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}

func GetCsv(srv *sheets.Service, sheetID string) ([][]interface{}, error) {
	resp, err := srv.Spreadsheets.Values.Get(sheetID, "シート1!A2:ZZ").Do()
	if err != nil {
		return nil, fmt.Errorf("get spread sheet data error: %v", err)
	}
	return resp.Values, nil
}

func (s *SheetClient) SheetID(sheetName string) (int64, error) {
	resp, err := s.srv.Spreadsheets.Get(s.spreadsheetID).Do()
	if err != nil {
		return 0, err
	}
	for _, sheet := range resp.Sheets {
		if sheet.Properties.Title == sheetName {
			return sheet.Properties.SheetId, nil
		}
	}
	return 0, fmt.Errorf("sheetName %s is not found", sheetName)
}

func (s *SheetClient) Get(range_ string) ([][]interface{}, error) {
	resp, err := s.srv.Spreadsheets.Values.Get(s.spreadsheetID, range_).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}


func GetAll() [][]interface{}{
	ctx := context.Background()
	client, err := NewSheetClient(ctx, "1LcGnzRuKYA5SyOBGsUm1oN1r4pNX-8lmaTYT54A_RsI")
	if err != nil {
		fmt.Println(err)
	}
	values, err := client.Get("'Answer'!A2:ZZ")
	return values
}
func GetCorporateSheet() [][]interface{}{
	ctx := context.Background()
	client, err := NewSheetClient(ctx, "1LcGnzRuKYA5SyOBGsUm1oN1r4pNX-8lmaTYT54A_RsI")
	if err != nil {
		fmt.Println(err)
	}
	values, err := client.Get("'Corporate'!A2:ZZ")
	return values
}
