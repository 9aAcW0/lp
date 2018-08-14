package modules

import (
	"encoding/base64"
	"strings"

	"io/ioutil"
	"log"

	"net/http"

	"fmt"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type PdfRequest struct {
	FileName    string
	FileContent string
	FilePassword string
}
type MailRequest struct {
	FileName    string
	FileContent string
	FilePassword string
}

func InitGmailClient() (client *http.Client, err error) {

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	client = getClient(config)
	fmt.Println("Done getClient!!")

	return client, err
}

func CreateMessageWithAttachment(client *http.Client, request MailRequest) (*gmail.Message, error) {

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
		return nil, err
	}

	var message gmail.Message

	// read file for attachment purpose
	// ported from https://developers.google.com/gmail/api/sendEmail.py

	//fileBytes, err := ioutil.ReadFile(request.FileName)
	//if err != nil {
	//	log.Fatalf("Unable to read file for attachment: %v", err)
	//}

	//fileMIMEType := http.DetectContentType(fileBytes)
	//
	//// https://www.socketloop.com/tutorials/golang-encode-image-to-base64-example
	////fileData := base64.StdEncoding.EncodeToString(fileBytes)
	//fileData := request.FileContent
	//
	//boundary := randStr(32, "alphanum")
	//
	//temp, err := toISO2022JP(
	//	"Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
	//
	//		"MIME-Version: 1.0\n" +
	//		"from: 'me'\n" +
	//		"reply-to:" + request.ReplyTo + "\n" +
	//		"to:" + request.To + "\n" +
	//		"subject: テストですよ\n\n" +
	//
	//		"--" + boundary + "\n" +
	//		request.Content + "\n\n" +
	//		"--" + boundary + "\n" +
	//
	//		"Content-Type: " + fileMIMEType + "; name=" + string('"') + request.FileName + string('"') + " \n" +
	//		"MIME-Version: 1.0\n" +
	//		"Content-Transfer-Encoding: base64\n" +
	//		"Content-Disposition: attachment; filename=" + string('"') + request.FileName + string('"') + " \n\n" +
	//		chunkSplit(fileData, 76, "\n") +
	//		"--" + boundary + "--")
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//// see https://godoc.org/google.golang.org/api/gmail/v1#Message on .Raw
	//// use URLEncoding here !! StdEncoding will be rejected by Google API
	//
	//message.Raw = base64.URLEncoding.EncodeToString(temp)

	msg, err := srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func CreateMessageWithNotAttachment() (*gmail.Message, error) {

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}
	client := getClient(config)
	spew.Dump("Done getClient!!")

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
		return nil, err
	}

	temp, err := toISO2022JP("From: 'me'\r\n" +
		"reply-to: ihukcl.9748@gmail.com\r\n" +
		"To: shimano@ispec.tech\r\n" +
		"Subject: テストですよ\r\n" +
		"\r\n" + "テストんご！")

	if err != nil {
		return nil, err
	}

	var message gmail.Message
	message.Raw = base64.StdEncoding.EncodeToString(temp)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "=", "", -1)

	msg, err := srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return nil, err
	}

	return msg, nil
}
