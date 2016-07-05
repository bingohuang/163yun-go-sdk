package cloudcomb

import (
	"bytes"
	"encoding/json"
	"errors"
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
	ExpiresIn uint64
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

/*=== user start 1 ===*/
// Get user's token
func (cc *CloudComb) PostUserToken() (string, uint64, error) {
	// user token request params
	type userTokenReq struct {
		AppKey    string `json:"app_key"`
		AppSecret string `json:"app_secret"`
	}
	utq := userTokenReq{
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
	type userTokenRes struct {
		Token     string `json:"token"`
		ExpiresIn uint64 `json:"expires_in"`
	}
	var uts userTokenRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&uts); err != nil {
		return "", 0, err
	}

	return uts.Token, uts.ExpiresIn, nil
}

/*=== user end ===*/

/*=== containers start 9 ===*/
// list all containers' images
func (cc *CloudComb) GetContainersImages() (string, error) {
	result, _, err := cc.doRESTRequest("GET", "/api/v1/containers/images", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// list all containers' info
func (cc *CloudComb) GetContainers() (string, error) {
	// TODO: limit=20&offset=0
	result, _, err := cc.doRESTRequest("GET", "/api/v1/containers", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// Get specified container's info
func (cc *CloudComb) GetContainer(id string) (string, error) {
	if id == "" {
		return "", errors.New("Container id is missed")
	}
	result, _, err := cc.doRESTRequest("GET", "/api/v1/containers/"+id, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// Get specified container's flow
func (cc *CloudComb) GetContainerFlow(id string) (string, error) {
	if id == "" {
		return "", errors.New("Container id is missed")
	}
	// TODO: from_time=1111&to_time=111111
	result, _, err := cc.doRESTRequest("GET", "/api/v1/containers/"+id+"/flow", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// TODO: create container


// TODO: update container

// TODO: delete container

// TODO: restart container

// TODO: tag a container to a image
/*=== containers end ===*/

/*=== clusters(apps) start 8 ===*/
// list all clusters' images
func (cc *CloudComb) GetClustersImages() (string, error) {
	result, _, err := cc.doRESTRequest("GET", "/api/v1/apps/images", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// list clusters
func (cc *CloudComb) GetClusters() (string, error) {
	// TODO: limit=20&offset=0
	result, _, err := cc.doRESTRequest("GET", "/api/v1/apps", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// get cluster
func (cc *CloudComb) GetCluster(id string) (string, error) {
	if id == "" {
		return "", errors.New("Cluster id is missed")
	}
	result, _, err := cc.doRESTRequest("GET", "/api/v1/apps/"+id, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// TODO: create cluster

// TODO: update cluster

// TODO: delete cluster

// TODO: replicate cluster

// TODO" watch cluster
/*=== clusters(apps) end ===*/

/*=== repositories start 4 ===*/
// list repositories
func (cc *CloudComb) GetRepositories() (string, error) {
	// TODO: limit=20&offset=0
	result, _, err := cc.doRESTRequest("GET", "/api/v1/repositories", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// get repository
func (cc *CloudComb) GetRepository(id string) (string, error) {
	if id == "" {
		return "", errors.New("Repository id is missed")
	}
	result, _, err := cc.doRESTRequest("GET", "/api/v1/repositories/"+id, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// TODO: create repository
// TODO: delete repository
/*=== repositories end ===*/

/*=== secret-keys start 4 ===*/

// list secret keys
func (cc *CloudComb) GetSecretKeys() (string, error) {
	result, _, err := cc.doRESTRequest("GET", "/api/v1/secret-keys", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// get secret key
func (cc *CloudComb) GetSecretKey(id string) (string, error) {
	if id == "" {
		return "", errors.New("Secret key id is missed")
	}
	result, _, err := cc.doRESTRequest("GET", "/api/v1/secret-keys/"+id, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// TODO: create secret key

// TODO: delete secret key
/*=== secret-keys end ===*/
