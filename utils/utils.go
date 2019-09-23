package utils

import "net/http"
import "net/url"
import "strings"

func BuildURLData(d *url.Values, k string, v string) {
	d.Set(k, v)
}

func MakeRequest(data url.Values) *http.Request {
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
