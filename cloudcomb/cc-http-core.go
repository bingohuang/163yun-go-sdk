package cloudcomb

import (
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
	"strings"
	"fmt"
	"io/ioutil"
	"errors"
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
	if method == "PUT" || method == "POST" {
		length := req.Header.Get("Content-Length")
		if length != "" {
			req.ContentLength, _ = strconv.ParseInt(length, 10, 64)
		}
	}

	return core.httpClient.Do(req)
}

func (cc *CloudComb) doRESTRequest(method, uri, query string, headers map[string]string,
value interface{}) (result string, rtHeaders http.Header, err error) {
	url := fmt.Sprintf("https://%s%s", cc.endpoint, uri)

	if query != "" {
		query = escapeURI(query)
		url += "?" + query
	}


	// GET and HEAD method have no body
	rc, ok := value.(io.Reader)
	if !ok || method == "GET" || method == "HEAD" {
		rc = nil
	}

	// header
	if headers == nil {
		headers = make(map[string]string)
	}
	// Content-Type:application/json
	headers["Content-Type"] = "application/json"

	// Normalize url
	if !strings.HasPrefix(uri, "/") {
		uri = "/" + uri
	}

	resp, err := cc.doHTTPRequest(method, url, headers, rc)
	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()

	// parse response
	// 20X
	if (resp.StatusCode / 100) == 2 {
		if method == "GET" && value != nil {
			written, err := chunkedCopy(value.(io.Writer), resp.Body)
			return strconv.FormatInt(written, 10), resp.Header, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		return string(body), resp.Header, err
	}
	// not 20X
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		if len(body) == 0 && (resp.StatusCode/100) != 2 {
			return "", resp.Header, errors.New(fmt.Sprint(resp.StatusCode))
		}
		return "", resp.Header, errors.New(string(body))
	} else {
		return "", resp.Header, err
	}

}

