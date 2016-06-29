package cloudcomb

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
)

// Cloudcomb Open API Client
type CloudComb struct {
	//core
	ccHTTPCore

	appKey    string
	appSecret string
	Token     string
	Expires   int
}

// Generate API token request payload
type UserToken struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func NewCC(appKey, appSecret string) *CloudComb {
	cc := &CloudComb{
		appKey:    appKey,
		appSecret: appSecret,
	}

	cc.httpClient = &http.Client{}
	cc.endpoint = defaultEndPoint
	cc.SetTimeout(defaultConnectTimeout)

	return cc
}

func (cc *CloudComb) getToken() (token string) {
	return "token " + cc.Token
}

func (cc *CloudComb) UserToken() (string, error) {
	ut := UserToken {
		AppKey: cc.appKey,
		AppSecret: cc.appSecret,
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(ut)

	result, _, err := cc.doRESTRequest("POST", "/api/v1/token", "", nil, body)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (cc *CloudComb) doRESTRequest(method, uri, query string, headers map[string]string,
	value interface{}) (result string, rtHeaders http.Header, err error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	// Normalize url
	if !strings.HasPrefix(uri, "/") {
		uri = "/" + uri
	}

	url := fmt.Sprintf("https://%s%s", cc.endpoint, uri)

	if query != "" {
		query = escapeURI(query)
		url += "?" + query
	}

	// header
	// Content-Type:application/json
	headers["Content-Type"] = "application/json"

	// GET and HEAD method have no body
	rc, ok := value.(io.Reader)
	if !ok || method == "GET" || method == "HEAD" {
		rc = nil
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
