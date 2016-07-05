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
	ccHTTPCore

	appKey    string
	appSecret string
	Token     string
	ExpiresIn uint64

	Container
}

type Container struct {
	ContainerID string
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
func (cc *CloudComb) UserToken() (string, uint64, error) {
	// user token request params
	type userTokenReq struct {
		AppKey    string `json:"app_key"`
		AppSecret string `json:"app_secret"`
	}
	req := userTokenReq{
		AppKey:    cc.appKey,
		AppSecret: cc.appSecret,
	}

	// generate json body
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(req)
	if err != nil {
		return "", 0, err
	}

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/token", "", nil, body)
	if err != nil {
		return "", 0, err
	}

	// user token response messages
	type userTokenRes struct {
		Token     string `json:"token"`
		ExpiresIn uint64 `json:"expires_in"`
	}
	var res userTokenRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return "", 0, err
	}

	return res.Token, res.ExpiresIn, nil
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

// create container
func (cc *CloudComb) CreateContainer(params string) (uint64, error) {
	if params == "" {
		return 0, errors.New("Params is missed")
	}
	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/containers", "", nil, body)
	if err != nil {
		return 0, err
	}

	// create container response messages
	type createContainerRes struct {
		Id uint64 `json:"id"`
	}
	var res createContainerRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return 0, err
	}

	return res.Id, nil
}

// update container
func (cc *CloudComb) UpdateContainer(id string, params string) error {
	if id == "" {
		return errors.New("Container id is missed")
	}
	if params == "" {
		return errors.New("Params is missed")
	}

	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	_, _, err := cc.doRESTRequest("PUT", "/api/v1/containers/"+id, "", nil, body)
	if err != nil {
		return err
	}
	return nil
}

// restart container
func (cc *CloudComb) RestartContainer(id string) error {
	if id == "" {
		return errors.New("Container id is missed")
	}
	// do rest request
	_, _, err := cc.doRESTRequest("PUT", "/api/v1/containers/"+id+"/actions/restart", "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// tag a container to a image
func (cc *CloudComb) TagContainer(id string, params string) (string, error) {
	if id == "" {
		return "", errors.New("Container id is missed")
	}
	if params == "" {
		return "", errors.New("Params is missed")
	}
	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/containers/"+id+"/tag", "", nil, body)
	if err != nil {
		return "", err
	}

	// create container response messages
	type tagContainerRes struct {
		ImageId string `json:"image_id"`
	}
	var res tagContainerRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return "", err
	}

	return res.ImageId, nil
}

// delete container
func (cc *CloudComb) DeleteContainer(id string) error {
	if id == "" {
		return errors.New("Container id is missed")
	}
	// do rest request
	_, _, err := cc.doRESTRequest("DELETE", "/api/v1/containers/"+id, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

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
