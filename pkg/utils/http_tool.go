package utils

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: time.Second * 5,
	}
}

func SendHttpRequest(method, url string, header map[string]string, bodyData []byte, maxRetry int) ([]byte, error) {

	resp := make([]byte, 0)
	log.Infof("send http request method:%v,url:%v,maxRetry:%v", method, url, maxRetry)
	log.Debugf("send http request method:%v,url:%v,header:%v,body:%v,maxRetry:%v", method, url, header, string(bodyData), maxRetry)
	body := &bytes.Buffer{}
	if bodyData != nil && len(bodyData) > 0 {
		body = bytes.NewBuffer(bodyData)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return resp, err
	}

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	// make sure maxRetry valid
	if maxRetry <= 0 {
		maxRetry = 1
	}

	sendReq := func(req *http.Request) ([]byte, error) {
		response, err := httpClient.Do(req)
		if err != nil {
			log.Errorf("send http request failed, err:%v", err)
			return nil, err
		}
		defer response.Body.Close()
		log.Infof("http request get response code:%v", response.StatusCode)
		data, err := ioutil.ReadAll(response.Body)
		if response.StatusCode != http.StatusOK {
			return nil, errors.New(fmt.Sprintf("http response code invalid,code:%v", response.StatusCode))
		}
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	for loop := 0; loop < maxRetry; loop++ {
		resp, err = sendReq(req)
		if err == nil {
			return resp, nil
		}
	}
	return resp, err
}
