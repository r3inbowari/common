package common

import (
	"encoding/json"
	"io"
	"net/http"
)

type RequestOptions struct {
	Method string
	Url    string
	http.Client
	Reader io.Reader
}

func RequestJson(options RequestOptions, v interface{}, interceptors ...func(reqPoint *http.Request, context *http.Client)) (*http.Response, error) {
	if options.Method == "" {
		options.Method = http.MethodGet
	}
	req, err := http.NewRequest(options.Method, options.Url, options.Reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	for _, interceptor := range interceptors {
		interceptor(req, &options.Client)
	}
	res, err := options.Do(req)

	if v != nil && res != nil && err == nil {
		err = json.NewDecoder(res.Body).Decode(&v)
	}
	return res, err
}
