package third_login

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

//基本配置
type Auth struct {
	AppId     string
	AppSecret string
	RedirUri  string
}

func HttpGet(url string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Transport: tr, Jar: cookieJar}

	res, err := client.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPost(url string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return "", err
	}
	response, _ := client.Do(res)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
