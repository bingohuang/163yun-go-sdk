package cloudcomb

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

// run go test by: CC_APPKEY="your app key" CC_APPSECRET="your app secret" go test -v
var (
	appKey    = os.Getenv("CC_APPKEY")
	appSecret = os.Getenv("CC_APPSECRET")
	cc        = NewCC(appKey, appSecret)
	repoName  = "openapi"
	tag       = "test"
	filePath  = "Dockerfile"
)

/*=== user start 1 ===*/
func TestCloudComb_UserToken(t *testing.T) {
	if token, expiresIn, err := cc.UserToken(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get token=%s and expires_in=%d\n\n", token, expiresIn)
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
		fmt.Printf("Get response: %s\n\n", res)
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
	if res, err := cc.CreateContainer(params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %d\n\n", res)
		cc.ContainerID = strconv.FormatUint(uint64(res), 10)
		time.Sleep(time.Second * 60)
	}
}

func TestCloudComb_GetContainers(t *testing.T) {
	if res, err := cc.GetContainers(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_GetContainer(t *testing.T) {
	if res, err := cc.GetContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_GetContainerFlow(t *testing.T) {
	if res, err := cc.GetContainerFlow(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
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
		fmt.Printf("Get response: %s\n\n", res)
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
		fmt.Printf("Update success. \n\n")
		time.Sleep(time.Second * 10)
	}
}

func TestCloudComb_RestartContainer(t *testing.T) {
	if err := cc.RestartContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Restart success. \n\n")
		time.Sleep(time.Second * 30)
	}
}

func TestCloudComb_DeleteContainer(t *testing.T) {
	if err := cc.DeleteContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Delete success. \n\n")
	}
}

/*=== clusters(apps) start 8 ===*/
func TestCloudComb_GetClustersImages(t *testing.T) {
	if res, err := cc.GetClustersImages(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_GetClusters(t *testing.T) {
	if res, err := cc.GetClusters(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_CreateCluster(t *testing.T) {
	params := `{
				"name": "%s",
				"image_type": %d,
				"image_id": %d,
				"spec_id": %d
			  }`
	params = fmt.Sprintf(params, "openapiCluster", 2, 38851, 1)
	if id, url, err := cc.CreateCluster(params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get Cluster: id=%d, url=%s\n\n", id, url)
		cc.ClusterID = strconv.FormatUint(uint64(id), 10)
		time.Sleep(time.Second * 60)
	}
}

func TestCloudComb_GetCluster(t *testing.T) {
	if res, err := cc.GetCluster(cc.ClusterID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_ReplicateCluster(t *testing.T) {
	if err := cc.ReplicateCluster(cc.ClusterID, 2); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Replicate success. \n\n")
		time.Sleep(time.Second * 30)
	}
}

func TestCloudComb_UpdateCluster(t *testing.T) {
	params := `{
	  "desc": "%s"
	}`
	params = fmt.Sprintf(params, fmt.Sprintf("Modify description: %v",  time.Now()))
	if err := cc.UpdateCluster(cc.ClusterID, params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Update success. \n\n")
		time.Sleep(time.Second *30)
	}
}

func TestCloudComb_WatchCluster(t *testing.T) {
	// use goroutine to watch the update result
	watch := make(chan bool, 1)
	go func() {

		if res, err := cc.WatchCluster(cc.ClusterID); err != nil {
			fmt.Println(err)
			t.Errorf("Fail to get response. %v", err)
		} else {
			fmt.Printf("Get watch response: %s\n\n", res)
		}
		watch <- true
	}()

	update := make(chan bool, 1)
	go func() {
		TestCloudComb_UpdateCluster(t)
		update <- true
	}()

	select {
	case <-watch: // 从watch中读取到数据
	case <-update: // 一直没有从watch中读取到数据,但从update中读取到了数据,这就超时结束了
	}
	time.Sleep(time.Second *10)
}

func TestCloudComb_DeleteCluster(t *testing.T) {
	if err := cc.DeleteCluster(cc.ClusterID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Delete success.\n\n")
	}
}

/*=== repositories start 4 ===*/
func TestCloudComb_GetRepositories(t *testing.T) {
	if res, err := cc.GetRepositories(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_CreateRepository(t *testing.T) {
	if err := cc.CreateRepository(repoName, tag, filePath); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Create success. \n\n")
	}
	time.Sleep(time.Second * 30)
}

func TestCloudComb_GetRepository(t *testing.T) {
	// can not get repository id from CreateRepository
	// get from another exist repository
	if res, err := cc.GetRepository("93"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_DeleteRepository(t *testing.T) {
	if err := cc.DeleteRepository(repoName, tag); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Delete success. \n\n")
	}
}

/*=== secret-keys start 4 ===*/
func TestCloudComb_GetSecretKeys(t *testing.T) {
	if res, err := cc.GetSecretKeys(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_CreateSecretKey(t *testing.T) {
	params := `{
				"key_name": "%s"
			  }`
	params = fmt.Sprintf(params, "OpenAPI")
	if id, name, err := cc.CreateSecretKey(params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get Secret Key: id=%d, name=%s\n\n", id, name)
		cc.SecretKeyID = strconv.FormatUint(uint64(id), 10)
		time.Sleep(time.Second * 1)
	}
}

func TestCloudComb_GetSecretKey(t *testing.T) {
	if res, err := cc.GetSecretKey(cc.SecretKeyID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_DeleteSecretKey(t *testing.T) {
	if err := cc.DeleteSecretKey(cc.SecretKeyID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Delete success. \n\n")
	}
}
