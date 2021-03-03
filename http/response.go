package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) (map[string]interface{}, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s", url), nil)
	if err != nil {
		return nil, fmt.Errorf("request: %s", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("client do: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode > 200 {
		return nil, fmt.Errorf("StatusCode : %d", response.StatusCode)
	}

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll : %s", err)
	}
	ret := make(map[string]interface{})
	err = json.Unmarshal(bs, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func Post(paras string, url string) (map[string]interface{}, error) {
	jsonStr := []byte(paras)
	ioBody := bytes.NewBuffer(jsonStr)
	request, err := http.NewRequest("POST", url, ioBody)
	if err != nil {
		return nil, fmt.Errorf("request %s", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do %s", err)
	}
	defer response.Body.Close()
	if response.StatusCode > 200 {
		return nil, fmt.Errorf("%d status code returned ,%s", response.StatusCode, url)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(bytes, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
