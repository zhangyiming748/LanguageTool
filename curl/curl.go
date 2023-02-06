package curl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func HttpPostJson(addHeaders map[string]string, data interface{}, urlPath string) (body []byte, err error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpGet(addHeaders map[string]string, data map[string]string, urlPath string) (body []byte, err error) {
	params := url.Values{}
	urlInfo, err := url.Parse(urlPath)
	if err != nil {
		log.Println(err)
	}
	for dataKey, dataVal := range data {
		params.Set(dataKey, dataVal)
	}
	urlInfo.RawQuery = params.Encode()
	fullUrl := urlInfo.String()
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
