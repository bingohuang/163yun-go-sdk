package cloudcomb

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
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

// New CloudComb
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

// Get user's token
func (cc *CloudComb) UserToken() (string, uint64, error) {
	// user token request params
	type UserTokenReq struct {
		AppKey    string `json:"app_key"`
		AppSecret string `json:"app_secret"`
	}
	utq := UserTokenReq{
		AppKey:    cc.appKey,
		AppSecret: cc.appSecret,
	}

	// generate json
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(utq)

	result, _, err := cc.doRESTRequest("POST", "/api/v1/token", "", nil, body)
	if err != nil {
		return "", 0, err
	}

	// user token response messages
	type UserTokenRes struct {
		Token     string `json:"token"`
		ExpiresIn uint64 `json:"expires_in"`
	}
	var uts UserTokenRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&uts); err != nil {
		return "", 0, err
	}

	return uts.Token, uts.ExpiresIn, nil
}

// List containers' images
func (cc *CloudComb) ContainersImages() (string, error) {
	result, _, err := cc.doRESTRequest("GET", "/api/v1/containers/images", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// List containers
func (cc *CloudComb) Containers() (string, error) {
	result, _, err := cc.doRESTRequest("GET", "/api/v1/containers", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}
