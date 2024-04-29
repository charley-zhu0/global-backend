/*
 * @Author: charley zhu
 * @Date: 2023-10-15 13:06:25
 * @LastEditTime: 2023-10-15 13:47:25
 * @LastEditors: charley zhu
 * @Description:
 */
package httpclient

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{
	Transport: tr,
}

func GetHttpClient() *http.Client {
	return client
}

func GET(url string, timeout time.Duration, headers map[string]string) (code int, resp string, err error) {
	return Factory("GET", url, timeout, nil, headers)
}

func POST(url string, body string, timeout time.Duration, headers map[string]string) (code int, resp string, err error) {
	reqBuf := strings.NewReader(body)
	return Factory("POST", url, timeout, reqBuf, headers)
}

func Factory(method string, url string, timeout time.Duration, body io.Reader, headers map[string]string) (code int, resp string, err error) {
	code = 500
	if len(url) == 0 {
		return code, resp, errors.New("url is empty")
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return code, resp, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	request = request.WithContext(ctx)

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	response, err := client.Do(request)
	if err != nil {
		return code, resp, err
	}
	// response may be nil
	defer response.Body.Close()

	code = response.StatusCode

	if body, err := io.ReadAll(response.Body); err != nil {
		return code, resp, err
	} else {
		return code, string(body), nil
	}
}
