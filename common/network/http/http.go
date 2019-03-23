package gghttp

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func DoHttpPost(url string, headers map[string]string, data url.Values, timeout time.Duration) ([]byte, error) {
	if len(url) == 0 {
		return nil, errors.New("url is empty")
	}

	c := &http.Client{}
	c.Timeout = time.Second * timeout
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DoHttpGet(url string, headers map[string]string, timeout time.Duration) ([]byte, error) {
	if len(url) == 0 {
		return nil, errors.New("url is empty")
	}

	c := &http.Client{}
	c.Timeout = time.Second * timeout

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
