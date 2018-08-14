package logserver

import (
	"bytes"
	"net/http"
	"openRPA-basic-module/models"

	"io/ioutil"

	json2 "encoding/json"

	"strconv"

	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin/json"
	"github.com/pkg/errors"
)

const (
	started    = "started"
	processing = "processing"
	succeeded  = "succeeded"
	failed     = "failed"
)

func Start(JSON *models.JSON, endPoint string) (log *models.Log, err error) {
	log, err = createLog(JSON, endPoint, started)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return log, nil
}
func Update(JSON *models.JSON, oldLog *models.Log) (log *models.Log, err error) {
	log, err = updateLog(JSON, oldLog, processing)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return log, nil
}
func Success(JSON *models.JSON, oldLog *models.Log) (log *models.Log, err error) {
	log, err = updateLog(JSON, oldLog, succeeded)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return log, nil
}
func Fail(JSON *models.JSON, oldLog *models.Log) (log *models.Log, err error) {
	log, err = updateLog(JSON, oldLog, failed)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return log, nil
}

func createLog(JSON *models.JSON, endPoint string, status string) (log *models.Log, err error) {
	url := "http://localhost:9001/tasks/"

	j := *JSON
	j.Input.Files = nil
	for _, v := range j.Exports {
		v.Files = nil
	}

	log = &models.Log{
		EndPoint: endPoint,
		Input:    j.Input,
		Exports:  j.Exports,
		Roots:    j.Roots,
		Status:   status,
	}

	resp, err := HttpPost(url, log)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var logResp *models.LogResponse
	err = json2.Unmarshal(body, &logResp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log = &logResp.Data.Task

	return log, nil
}

func updateLog(JSON *models.JSON, oldLog *models.Log, status string) (log *models.Log, err error) {
	url := "http://localhost:9001/tasks/"

	j := *JSON
	j.Input.Files = nil
	for _, v := range j.Exports {
		v.Files = nil
	}
	j.Roots = []models.Root{}

	log = &models.Log{
		Input:    j.Input,
		Exports:  j.Exports,
		Roots:    j.Roots,
		Status:   status,
		Model:    oldLog.Model,
		EndPoint: oldLog.EndPoint,
	}

	resp, err := HttpPut(url, oldLog.ID, log)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	spew.Dump(oldLog.ID)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("ログサーバとの接続に失敗しました。StatusCode:" + strconv.Itoa(resp.StatusCode))
	}

	err = json2.Unmarshal(body, &log)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return log, nil
}

func HttpPost(url string, body interface{}) (*http.Response, error) {

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		return nil, err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func HttpPut(url string, id uint64, body interface{}) (*http.Response, error) {

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		url+strconv.FormatUint(id, 10),
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		return nil, err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
