package cloudcomb

import (
	"fmt"
	"os"
	"testing"
	"time"
	"strconv"
)

// run go test by: CC_APPKEY="your app key" CC_APPSECRET="your app secret" go test -v
var (
	appKey    = os.Getenv("CC_APPKEY")
	appSecret = os.Getenv("CC_APPSECRET")
	cc        = NewCC(appKey, appSecret)
)

/*=== user start 1 ===*/
func TestCloudComb_UserToken(t *testing.T) {
	if token, expiresIn, err := cc.UserToken(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get token=%s and expires_in=%d\n", token, expiresIn)
		cc.Token = token
	}
}

/*=== containers start 9 ===*/
// https://c.163.com/wiki/index.php?title=容器API
func TestCloudComb_GetContainersImages(t *testing.T) {
	if res, err := cc.GetContainersImages(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_CreateContainer(t *testing.T) {
	params := `{
				"image_type": %d,
				"image_id": %d,
				"name": "%s",
				"desc": "%s",
				"ssh_key_ids": %s,
				"env_var": {
					"key1": "value1",
    				"key2": "value2"
				},
				"charge_type": 2,
				"spec_id": 1,
				"use_public_network": 1,
				"network_charge_type": 1,
				"bandwidth": 1
			  }`
	params = fmt.Sprintf(params, 1, 20835, "openapitest", "cloudcomb open api test container", "[]")
	fmt.Printf("params: %s\n", params)
	if res, err := cc.CreateContainer(params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %d\n", res)
		cc.ContainerID = strconv.FormatUint(res, 10)
		time.Sleep(time.Second * 60)
	}
}

func TestCloudComb_GetContainers(t *testing.T) {
	if res, err := cc.GetContainers(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_GetContainer(t *testing.T) {
	if res, err := cc.GetContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_GetContainerFlow(t *testing.T) {
	if res, err := cc.GetContainerFlow(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_TagContainer(t *testing.T) {
	params := `{
	  "repo_name": "%s",
	  "tag": "%s"
	}`
	fmt.Printf("params: %s\n", params)
	params = fmt.Sprintf(params, "openapi", time.Now().Format("20060102150405"))
	if res, err := cc.TagContainer(cc.ContainerID, params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_UpdateContainer(t *testing.T) {
	params := `{
	  "charge_type":%d,
	  "desc": "%s",
	  "network_charge_type":%d,
	  "bandwidth":%d
	}`
	params = fmt.Sprintf(params, 2, "Modify description", 1, 2)
	fmt.Printf("params: %s\n", params)
	if err := cc.UpdateContainer(cc.ContainerID, params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Update success\n")
		time.Sleep(time.Second * 10)
	}
}

func TestCloudComb_RestartContainer(t *testing.T) {
	if err := cc.RestartContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Restart success\n")
		time.Sleep(time.Second * 30)
	}
}

func TestCloudComb_DeleteContainer(t *testing.T) {
	if err := cc.DeleteContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Delete success\n")
	}
}

/*=== clusters(apps) start 8 ===*/
func TestCloudComb_GetClustersImages(t *testing.T) {
	if res, err := cc.GetClustersImages(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_GetClusters(t *testing.T) {
	if res, err := cc.GetClusters(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_GetCluster(t *testing.T) {
	if res, err := cc.GetCluster("413529"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

/*=== repositories start 4 ===*/
func TestCloudComb_GetRepositories(t *testing.T) {
	if res, err := cc.GetRepositories(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_GetRepository(t *testing.T) {
	if res, err := cc.GetRepository("22103"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

/*=== secret-keys start 4 ===*/
func TestCloudComb_GetSecretKeys(t *testing.T) {
	if res, err := cc.GetSecretKeys(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_GetSecretKey(t *testing.T) {
	if res, err := cc.GetSecretKey("196"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}
