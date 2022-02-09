package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PostJson(url string, data interface{}) (body string, err error) {
	b, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
