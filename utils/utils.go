package main

import "net/http"
import "net/url"
import "strings"

func buildURLData(d *url.Values, k string, v string) {
	d.Set(k, v)
}

func makeRequest(data url.Values) *http.Request {
	req, _ := http.NewRequest("POST", runningConfig.Host+runningConfig.Endpoint, strings.NewReader(data.Encode()))
	for _, header := range runningConfig.Headers {
		if header.Type == "Cookie" {
			cookie := http.Cookie{Name: header.Name, Value: header.Value}
			req.AddCookie(&cookie)
		} else {
			req.Header.Set(header.Type, header.Value)
		}

	}
	return req
}
