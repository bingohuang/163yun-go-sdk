package cloudcomb

import (
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

// cloudcomb http core struct
type ccHTTPCore struct {
	endpoint   string
	httpClient *http.Client
}

// set timeout
func (core *ccHTTPCore) SetTimeout(timeout int) {
	core.httpClient = &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (c net.Conn, err error) {
				c, err = net.DialTimeout(network, addr, time.Duration(timeout)*time.Second)
				if err != nil {
					return nil, err
				}
				return
			},
			// http://studygolang.com/articles/3138
			// DisableKeepAlives: true,
		},
	}
}

// do http request
func (core *ccHTTPCore) doHTTPRequest(method, url string, headers map[string]string,
	body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// User Agent
	req.Header.Set("User-Agent", makeUserAgent())

	// https://code.google.com/p/go/issues/detail?id=6738
	if method == "PUT" || method == "POSt" {
		length := req.Header.Get("Content-Length")
		if length != "" {
			req.ContentLength, _ = strconv.ParseInt(length, 10, 64)
		}
	}

	return core.httpClient.Do(req)
}
