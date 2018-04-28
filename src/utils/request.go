/**
 * Created by chaolinding on 2018/4/8.
 */

package utils

import (
	"net/http"
	"strings"
	"time"
	"fmt"
	"bytes"
)

/*
POST 请求
 */
func Post(url string, postData map[string]string) (int, error) {

	client := &http.Client{Timeout: time.Second * 60}
	body := "{"
	length := len(postData)
	count := 0
	for key, value := range postData {
		count ++
		body += "\"" + key + "\":\"" + value + "\""
		if count < length {
			body += ","
		}

	}
	body += "}"
	fmt.Println(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(([]byte(body))))
	if err != nil {
		return 400, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "timeout") {
			return 408, err
		}
		return 400, err
	}
	return resp.StatusCode, nil
}

/*
GET 请求
 */
func get(url string) (int, error) {

	resp, err := http.Get(url)

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "timeout") {
			return 408, err
		}
		return 400, err
	}
	return resp.StatusCode, nil
}
