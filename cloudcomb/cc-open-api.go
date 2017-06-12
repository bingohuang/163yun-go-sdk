package cloudcomb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Cloudcomb Open API Client
type CloudComb struct {
	ccHTTPCore

	appKey    string
	appSecret string

	User

	Container

	Cluster

	Repository

	SecretKey

	Namespace

	Microservice

	Ips
}

type User struct {
	Token string
}

type Container struct {
	ContainerID string
}

type Cluster struct {
	ClusterID string
}

type Repository struct {
	RepositoryID string
}

type SecretKey struct {
	SecretKeyID string
}

type Namespace struct {
	NamespaceID string
}

type Microservice struct {
	ServiceID string
}

type Ips struct {
	IpIDs []string
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

/*=== user start count=1 ===*/
// Get user's token
func (cc *CloudComb) UserToken() (string, error) {
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
		return "", err
	}

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/token", "", nil, body)
	if err != nil {
		return "", err
	}

	// user token response messages
	type userTokenRes struct {
		Token string `json:"token"`
		//ExpiresIn uint   `json:"expires_in"`
	}
	var res userTokenRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return "", err
	}

	return res.Token, nil
}

/*=== user end ===*/

/*=== containers start count=9 ===*/
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
	// TODO query: limit=20&offset=0
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
	// TODO query: from_time=1111&to_time=111111
	result, _, err := cc.doRESTRequest("GET", "/api/v1/containers/"+id+"/flow", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// create container
func (cc *CloudComb) CreateContainer(params string) (uint, error) {
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
		Id uint `json:"id"`
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

/*=== clusters(apps) start count=8 ===*/
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
	// TODO query: limit=20&offset=0
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

// create cluster
func (cc *CloudComb) CreateCluster(params string) (uint, string, error) {
	if params == "" {
		return 0, "", errors.New("Params is missed")
	}
	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/apps", "", nil, body)
	if err != nil {
		return 0, "", err
	}

	// create cluster response messages
	type createClusterRes struct {
		Id  uint   `json:"id"`
		Url string `json:"url"`
	}
	var res createClusterRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return 0, "", err
	}

	return res.Id, res.Url, nil
}

// update cluster
func (cc *CloudComb) UpdateCluster(id string, params string) error {
	if id == "" {
		return errors.New("Cluster id is missed")
	}
	if params == "" {
		return errors.New("Params is missed")
	}

	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	_, _, err := cc.doRESTRequest("PUT", "/api/v1/apps/"+id, "", nil, body)
	if err != nil {
		return err
	}
	return nil
}

// replicate cluster
func (cc *CloudComb) ReplicateCluster(id string, replicas int) error {
	if id == "" {
		return errors.New("Cluster id is missed")
	}
	// do rest request
	_, _, err := cc.doRESTRequest("PUT", "/api/v1/apps/"+id+"/replications/"+strconv.Itoa(replicas)+"/actions/resize", "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// watch cluster
func (cc *CloudComb) WatchCluster(id string) (string, error) {
	if id == "" {
		return "", errors.New("Cluster id is missed")
	}
	// support long connection
	result, _, err := cc.doRESTRequest("GET", "/api/v1/watch/apps/"+id, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// delete cluster
func (cc *CloudComb) DeleteCluster(id string) error {
	if id == "" {
		return errors.New("Cluster id is missed")
	}
	// do rest request
	_, _, err := cc.doRESTRequest("DELETE", "/api/v1/apps/"+id, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

/*=== clusters(apps) end ===*/

/*=== repositories start count=4 ===*/
// list repositories
func (cc *CloudComb) GetRepositories() (string, error) {
	// TODO query: limit=20&offset=0
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

// create repository
func (cc *CloudComb) CreateRepository(repoName string, tag string, path string) error {
	if repoName == "" {
		return errors.New("Repository repoName is missed")
	}
	if tag == "" {
		return errors.New("Repository tag is missed")
	}
	if path == "" {
		return errors.New("Repository file path is missed")
	}

	_, _, err := cc.doFormRequest("/api/v1/repositories/"+repoName+"/tags/"+tag+"/actions/build", nil, "docker_file", path)
	if err != nil {
		return err
	}

	return nil
}

// delete repository
func (cc *CloudComb) DeleteRepository(repoName string, tag string) error {
	if repoName == "" {
		return errors.New("Repository repoName is missed")
	}
	if tag == "" {
		return errors.New("Repository tag is missed")
	}
	// do rest request
	_, _, err := cc.doRESTRequest("DELETE", "/api/v1/repositories/"+repoName+"/tags/"+tag, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

/*=== repositories end ===*/

/*=== secret-keys start count=4 ===*/

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

// create secret key
func (cc *CloudComb) CreateSecretKey(params string) (uint, string, error) {
	if params == "" {
		return 0, "", errors.New("Params is missed")
	}
	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/secret-keys", "", nil, body)
	if err != nil {
		return 0, "", err
	}

	// create cluster response messages
	type createSecretKeyRes struct {
		Id          uint   `json:"id"`
		Name        string `json:"name"`
		FingerPrint string `json:"fingerprint"`
		CreatedAt   string `json:"created_at"`
	}
	var res createSecretKeyRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return 0, "", err
	}

	return res.Id, res.Name, nil
}

// delete secret key
func (cc *CloudComb) DeleteSecretKey(id string) error {
	if id == "" {
		return errors.New("Secret key id is missed")
	}
	// do rest request
	_, _, err := cc.doRESTRequest("DELETE", "/api/v1/secret-keys/"+id, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

/*=== secret-keys end ===*/

/*=== namespaces start count=4 ===*/

// create namespace
func (cc *CloudComb) CreateNamespace(params string) (uint, error) {
	if params == "" {
		return 0, errors.New("Params is missed")
	}
	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/namespaces", "", nil, body)
	if err != nil {
		return 0, err
	}

	// create cluster response messages
	type createNamespaceRes struct {
		Id uint `json:"namespace_id"`
	}
	var res createNamespaceRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return 0, err
	}

	return res.Id, nil
}

// list namespaces
func (cc *CloudComb) GetNamespaces() (string, error) {
	result, _, err := cc.doRESTRequest("GET", "/api/v1/namespaces", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// get namespace services
func (cc *CloudComb) GetNamespaceServices(id string, offset int, limit int) (string, error) {
	if id == "" {
		return "", errors.New("Namespace id is missed")
	}
	uri := "/api/v1/namespaces/" + id + "/microservices"
	// TODO: ?offset=0&limit=10
	//if offset >= 0 || limit > 0 {
	//	uri += "?"
	//}
	result, _, err := cc.doRESTRequest("GET", uri, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// delete namespace
func (cc *CloudComb) DeleteNamespace(id string) error {
	if id == "" {
		return errors.New("Namespace id is missed")
	}
	// do rest request
	_, _, err := cc.doRESTRequest("DELETE", "/api/v1/namespaces/"+id, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

/*=== namespaces end ===*/

/*=== microservices start count= ===*/

// create microservice
func (cc *CloudComb) CreateMicroservice(params string) (uint, error) {
	if params == "" {
		return 0, errors.New("Params is missed")
	}
	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/microservices", "", nil, body)
	if err != nil {
		return 0, err
	}

	// create cluster response messages
	type createMicroserviceRes struct {
		Id uint `json:"service_id"`
	}
	var res createMicroserviceRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return 0, err
	}

	return res.Id, nil
}

// get microservice
func (cc *CloudComb) GetMicroservice() (string, error) {
	result, _, err := cc.doRESTRequest("GET", "/api/v1/microservices", "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// delete microservice
func (cc *CloudComb) DeleteMicroservice(id string, freeIp bool) error {
	if id == "" {
		return errors.New("Microservice id is missed")
	}
	// do rest request
	uri := "/api/v1/microservices/" + id + "?free_ip=" + strconv.FormatBool(freeIp)
	_, _, err := cc.doRESTRequest("DELETE", uri, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

/*=== microservices end ===*/

/*=== IP start count=6 ===*/

// create IP
func (cc *CloudComb) CreateIP(params string) ([]string, []string, error) {
	if params == "" {
		return nil, nil, errors.New("Params is missed")
	}
	params = PurifyParams(params)

	body := bytes.NewBufferString(params)

	// do rest request
	result, _, err := cc.doRESTRequest("POST", "/api/v1/ips", "", nil, body)
	if err != nil {
		return nil, nil, err
	}

	// create response messages

	type ip struct {
		Id          string `json:"id"`
		Ip          string `json:"ip"`
		Status      string `json:"status"`
		IpType      string `json:"type"`
		ServiceId   string `json:"service_id"`
		ServiceName string `json:"service_name"`
		ServiceType string `json:"service_type"`
		CreateAt    string `json:"create_at"`
		UpdateAt    string `json:"update_at"`
	}

	type createIpRes struct {
		Total uint `json:"total"`
		Ips   []ip `json:"ips"`
	}
	var res createIpRes

	// parse json
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&res); err != nil {
		return nil, nil, err
	}

	var ids []string
	var ips []string
	for _, v := range res.Ips {
		ids = append(ids, v.Id)
		ips = append(ips, v.Ip)
	}

	return ids, ips, nil
}

// list IPs
func (cc *CloudComb) GetIPs(status string, iptype string, offset int, limit int) (string, error) {
	uri := "/api/v1/ips?" + fmt.Sprint(offset) + "&" + fmt.Sprint(limit)
	if status != "" {
		uri += "&" + status
	}
	if iptype != "" {
		uri += "&" + iptype
	}

	result, _, err := cc.doRESTRequest("GET", uri, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// get IP
func (cc *CloudComb) GetIP(id string) (string, error) {
	if id == "" {
		return "", errors.New("IP id is missed")
	}
	uri := "/api/v1/ips/" + id

	result, _, err := cc.doRESTRequest("GET", uri, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// get IP quota
func (cc *CloudComb) GetIPQuota() (string, error) {
	uri := "api/v1/ips/quota"

	result, _, err := cc.doRESTRequest("GET", uri, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// get IP cost
func (cc *CloudComb) GetIPCost(id string) (string, error) {
	if id == "" {
		return "", errors.New("IP id is missed")
	}
	uri := "api/v1/ips/" + id + "/cost"

	result, _, err := cc.doRESTRequest("GET", uri, "", nil, nil)
	if err != nil {
		return "", err
	}

	return result, nil
}

// delete IP
func (cc *CloudComb) DeleteIP(id string) error {
	if id == "" {
		return errors.New("IP id is missed")
	}
	uri := "api/v1/ips/" + id
	// do rest request
	_, _, err := cc.doRESTRequest("DELETE", uri, "", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

/*=== IP end ===*/
