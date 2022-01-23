package request

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"net/http"
)

// HTTPPost 发起一个 http post 请求
func HTTPPost(url string, data io.Reader, headers map[string]string) (body []byte, err error) {
	body = []byte{}
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Close = true
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ = ioutil.ReadAll(resp.Body)
		erroMsg := fmt.Sprintf("HTTP Connect Failed, Code is %d, body is %s", resp.StatusCode, string(body))
		err = errors.New(erroMsg)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// HTTPGet 发起一个 http get 请求
func HTTPGet(url string, data io.Reader, headers map[string]string) (body []byte, err error) {
	body = []byte{}
	req, err := http.NewRequest("GET", url, data)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Close = true
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ = ioutil.ReadAll(resp.Body)
		erroMsg := fmt.Sprintf("HTTP Connect Failed, Code is %d, body is %s", resp.StatusCode, string(body))
		err = errors.New(erroMsg)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// HTTPPut 发起一个 http put 请求
func HTTPPut(url string, data io.Reader, headers map[string]string) (body []byte, err error) {
	body = []byte{}
	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Close = true
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ = ioutil.ReadAll(resp.Body)
		erroMsg := fmt.Sprintf("HTTP Connect Failed, Code is %d, body is %s", resp.StatusCode, string(body))
		err = errors.New(erroMsg)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// HTTPDelete 发起一个 http delete 请求
func HTTPDelete(url string, data io.Reader, headers map[string]string) (body []byte, err error) {
	body = []byte{}
	req, err := http.NewRequest("DELETE", url, data)
	if err != nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Close = true
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ = ioutil.ReadAll(resp.Body)
		erroMsg := fmt.Sprintf("HTTP Connect Failed, Code is %d, body is %s", resp.StatusCode, string(body))
		err = errors.New(erroMsg)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
