package api

import (
	"openRPA-pdf-module-go/app/manager"
	"openRPA-basic-module/models"
	"openRPA-pdf-module-go/app/modules"
	"openRPA-pdf-module-go/app/validators"
	"time"
	"openRPA-pdf-module-go/app/manager/log_server"
	"openRPA-basic-module/api"
	"openRPA-pdf-module-go/app/lib/tool"

	"google.golang.org/api/gmail/v1"
	"os/exec"
)

type Gmail struct{}

func (ctrl Gmail) Post(c openrpabaseapi.Api) *openrpabaseapi.Api {

	var err error

	if err := c.BindJSON(&c.JSON); err != nil {
		return c.SetMessage("Faild Bind JSON error message:" + err.Error()).HandleBadRequestError()
	}
	hostUrl, err := tool.GetHostURL()
	if err != nil {
		return c.SetMessage("Hostの取得に失敗しました").SetErrorCode(2222).HandleInternalServerError()
	}

	// エンドポイントを指定
	c.EndPoint = "https://rpa.ispec.tech/gmail/post"
	// このエンドポイント用のexportの空モデルを追加する
	c.JSON.Exports = append(c.JSON.Exports, models.Export{})

	// ログサーバに新しいモデルをpostする
	//c.Log, err = logserver.Start(c.JSON, c.EndPoint)
	if err != nil {
		return c.SetMessage("Logサーバへのpostへ失敗しました エラー:" + err.Error()).HandleInternalServerError()
	}
	c.JSON.TaskLineID = c.Log.TaskLineID
	myTaskIndex := manager.EndPointIndex(c.JSON, c.EndPoint)
	if myTaskIndex == 0 {
		return c.RenderJSON(nil)
	}

	if err := validators.Struct(c.JSON); err != nil {
		return c.SetMessage(err.Error()).HandleBadRequestError()
	}

	if err := manager.RetrieveDataFromPointer(c.JSON, c.EndPoint); err != nil {
		return c.SetMessage(err.Error()).HandleInternalServerError()
	}

	mailAddresses, err := manager.RetrieveArgumentFromContents(c.JSON, c.EndPoint, "mail_address")
	if err != nil {
		return c.SetMessage("Emailアドレス一覧の取得に失敗しました エラー：" + err.Error()).SetErrorCode(1212).HandleBadRequestError()
	}

	pdfNames, err := manager.RetrieveArgumentFromContents(c.JSON, c.EndPoint, "pdf_name")
	if err != nil {
		return c.SetMessage("pdf_name一覧の取得に失敗しました エラー：" + err.Error()).SetErrorCode(1212).HandleBadRequestError()
	}

	pdfContents, err := manager.RetrieveArgumentFromContents(c.JSON, c.EndPoint, "pdf_content")
	if err != nil {
		return c.SetMessage("pdf_content一覧の取得に失敗しました エラー：" + err.Error()).SetErrorCode(1212).HandleBadRequestError()
	}
	//passつけるよん
	pdfPasswords, err := manager.RetrieveArgumentFromContents(c.JSON, c.EndPoint, "pdf_password")
	if err != nil {
		return c.SetMessage("pdf_password一覧の取得に失敗しました エラー：" + err.Error()).SetErrorCode(1212).HandleBadRequestError()
	}

	logserver.Update(c.JSON, c.Log)

	mailLen := len(mailAddresses)
	mailRequests := make([]modules.MailRequest, mailLen)

	var results []*gmail.Message

	cli, err := modules.InitGmailClient()
	if err != nil {
		return c.SetMessage("Gmailとの接続に失敗しました エラー：" + err.Error()).SetErrorCode(1111).HandleInternalServerError()
	}

	var export models.Export
	for i := range mailRequests {

		err := exec.Command("pdftk input.pdf output %s user_pw %s",pdfNames[i],pdfPasswords[i]).Run()
		if err != nil {
			return c.SetMessage("PDFのpasswordの設定に失敗しました　エラー：" + err.Error()).SetErrorCode(1111).HandleInternalServerError()
		}

		mailRequest := modules.MailRequest{
			FileName:    pdfNames[i],
			FileContent: pdfContents[i],
			FilePassword: pdfPasswords[i],
		}

		mailRequests[i] = mailRequest
		result, err := modules.CreateMessageWithAttachment(cli, mailRequest)
		if err != nil {
			return c.SetMessage("テストメール項目に失敗しました エラー：" + err.Error()).SetErrorCode(1111).HandleInternalServerError()
		}

		results = append(results, result)

		export = models.Export{
			ExecDate:     time.Now().String(),
			FromEndPoint: hostUrl,
			Data: struct {
				Results []*gmail.Message `json:"results"`
			}{results},
		}
		c.JSON.Exports[len(c.JSON.Exports)-1] = export
		logserver.Update(c.JSON, c.Log)
	}

	return c.RenderJSON(nil)
}
