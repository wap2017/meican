package api

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (mc *MeiCan) buildUrl(partUrl string, paramMap map[string]string) string {
	url := mc.baseUrl + partUrl + "?"
	for k, v := range paramMap {
		url += k + "=" + v + "&"
	}
	for k, v := range mc.baseParams {
		url += k + "=" + v + "&"
	}
	if strings.HasSuffix(url, "&") {
		url = url[:len(url)-1]
	}
	//log.Printf("url=%v", url)
	return url
}

func (mc *MeiCan) Get(partUrl string, m map[string]string) (*http.Response, error) {
	cli := http.Client{}

	req, err := http.NewRequest(http.MethodGet,
		mc.buildUrl(partUrl, m), nil)
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	for _, cookie := range mc.cookies {
		req.AddCookie(cookie)
	}

	return cli.Do(req)
}

func (mc *MeiCan) Post(partUrl string, m map[string]string) (*http.Response, error) {
	cli := http.Client{}

	data := url.Values{}
	for k, v := range m {
		data.Set(k, v)
	}

	//log.Printf("url=%v", mc.buildUrl(partUrl, make(map[string]string)))
	req, err := http.NewRequest(http.MethodPost,
		mc.buildUrl(partUrl, nil), strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	for _, cookie := range mc.cookies {
		req.AddCookie(cookie)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return cli.Do(req)
}


