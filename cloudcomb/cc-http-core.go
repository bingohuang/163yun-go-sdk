package cloudcomb

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

// purify params by "\n", "\t"
func PurifyParams(params string) string {
	params = strings.Replace(params, "\n", "", -1)
	params = strings.Replace(params, "\t", "", -1)
	return params
}

// do http request
func (core *ccHTTPCore) doHTTPRequest(method, url string, headers map[string]string,
	body io.Reader) (*http.Response, error) {
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
	value interface{}) (string, http.Header, error) {
	// Normalize url
	if !strings.HasPrefix(uri, "/") {
		uri = "/" + uri
	}
	url := fmt.Sprintf("https://%s%s", cc.endpoint, uri)
	//fmt.Printf("url=%s\n", url)

	// Normalize query
	if query != "" {
		query = escapeURI(query)
		url += "?" + query
	}

	// Normalize header
	if headers == nil {
		headers = make(map[string]string)
	}
	if cc.Token != "" {
		// Authorization:Token xxxxxxxxxxxxxx
		headers["Authorization"] = "Token " + cc.Token

	}
	// Content-Type:application/json
	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "application/json"
	}
	//fmt.Printf("headers=%v\n", headers)

	// body
	rc, ok := value.(io.Reader)
	// GET and HEAD method have no body
	if !ok || method == "GET" || method == "HEAD" {
		rc = nil
	}
	if rc != nil {
		//fmt.Printf("body=%s\n", rc)
	}

	// do HTTP request
	resp, err := cc.doHTTPRequest(method, url, headers, rc)
	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()

	// parse HTTP response
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

func (cc *CloudComb) doFormRequest(url string, params map[string]string, paramName, path string) (string, http.Header, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{} // generate form data
	headers := make(map[string]string)

	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("docker_file", filepath.Base(path))
	if err != nil {
		return "", nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", nil, err
	}

	headers["Content-Type"] = writer.FormDataContentType()

	err = writer.Close()

	if err != nil {
		return "", nil, err
	}

	// do rest request
	return cc.doRESTRequest("POST", url, "", headers, body)
}
