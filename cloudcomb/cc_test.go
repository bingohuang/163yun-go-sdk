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

/*=== user start count=1 ===*/
func TestCloudComb_UserToken(t *testing.T) {
	if token, err := cc.UserToken(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get token=%s\n\n", token)
		cc.Token = token
	}
}

/*=== containers start count=9 ===*/
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
		fmt.Print("Update success. \n\n")
		time.Sleep(time.Second * 10)
	}
}

func TestCloudComb_RestartContainer(t *testing.T) {
	if err := cc.RestartContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Restart success. \n\n")
		time.Sleep(time.Second * 30)
	}
}

func TestCloudComb_DeleteContainer(t *testing.T) {
	if err := cc.DeleteContainer(cc.ContainerID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Delete success. \n\n")
	}
}

/*=== clusters(apps) start count=8 ===*/
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
		fmt.Print("Replicate success. \n\n")
		time.Sleep(time.Second * 30)
	}
}

func TestCloudComb_UpdateCluster(t *testing.T) {
	params := `{
	  "desc": "%s"
	}`
	params = fmt.Sprintf(params, fmt.Sprintf("Modify description: %v", time.Now()))
	if err := cc.UpdateCluster(cc.ClusterID, params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Update success. \n\n")
		time.Sleep(time.Second * 30)
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
	time.Sleep(time.Second * 10)
}

func TestCloudComb_DeleteCluster(t *testing.T) {
	if err := cc.DeleteCluster(cc.ClusterID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Delete success.\n\n")
	}
}

/*=== repositories start count=4 ===*/
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
		fmt.Print("Create success. \n\n")
	}
	time.Sleep(time.Second * 60)
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
		fmt.Print("Delete success. \n\n")
	}
}

/*=== secret-keys start count=4 ===*/
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
		fmt.Print("Delete success. \n\n")
	}
}

/*=== namespaces start count=5 ===*/
func TestCloudComb_CreateNamespace(t *testing.T) {
	params := `{
				"name": "%s"
			  }`
	params = fmt.Sprintf(params, "test-namespace")
	if id, err := cc.CreateNamespace(params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get Namespace: id=%d\n\n", id)
		cc.NamespaceID = fmt.Sprint(id)
		time.Sleep(time.Second * 1)
	}
}

func TestCloudComb_GetNamespaces(t *testing.T) {
	if res, err := cc.GetNamespaces(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_GetNamespaceServices(t *testing.T) {
	if res, err := cc.GetNamespaceServices(cc.NamespaceID, -1, -1); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_DeleteNamespace(t *testing.T) {
	if err := cc.DeleteNamespace(cc.NamespaceID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Delete success. \n\n")
	}
}

/*=== microservices start count=3 ===*/
func TestCloudComb_CreateMicroservice(t *testing.T) {
	params := `{
		"bill_info":"default",
		"service_info": {
			"namespace_id": "93210",
			"stateful": 1,
			"replicas": 1,
			"service_name": "test-ubuntu",
			"port_maps": [
				{
					"target_port": "80",
					"port": "8080",
					"protocol": "TCP"
				}
			],
			"spec_alias": "C1M2S20",
			"state_public_net": {
				"used": true,
				"type": "flow",
				"bandwidth": 20
			},
			"disk_type": 2,
			"ip_id": "3e90c24d-2f0b-4fc8-a0c1-53220259dc54"
		},
		"service_container_infos": [
			{
				"image_path": "hub.c.163.com/public/ubuntu:14.04",
				"container_name": "test-ubuntu",
				"command": "",
				"envs": [
					{
						"key": "password",
						"value": "password"
					},
					{
						"key": "user",
						"value": "user"
					}
				],
				"log_dirs": [
					"/var/log/"
				],
				"cpu_weight": 100,
				"memory_weight": 100,
				"ssh_keys": [
					"test"
				]
			}
		]
	}`
	// create stateful service
	//params = fmt.Sprintf(params, "test-namespace")
	if id, err := cc.CreateMicroservice(params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get Microservice: id=%d\n\n", id)
		cc.ServiceID = fmt.Sprint(id)
		time.Sleep(time.Second * 30)
	}
}

func TestCloudComb_GetMicroservice(t *testing.T) {
	if res, err := cc.GetMicroservice(cc.ServiceID); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_DeleteMicroservice(t *testing.T) {
	// wait microservice state
	time.Sleep(time.Second * 30)
	if err := cc.DeleteMicroservice(cc.ServiceID, false); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Delete success. \n\n")
	}
}

/*=== IP start count=4 ===*/
func TestCloudComb_CreateIP(t *testing.T) {
	params := `{
				"nce": "%d",
				"nlb": "%d"
			  }`
	params = fmt.Sprintf(params, 1, 1)
	if ids, ips, err := cc.CreateIP(params); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get IP: ids=%v\n\n", ids)
		fmt.Printf("Get IP: ips=%v\n\n", ips)
		cc.IpIDs = ids
		time.Sleep(time.Second * 1)
	}
}

func TestCloudComb_GetIPs(t *testing.T) {
	if res, err := cc.GetIPs("", "", 0, 2); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_GetIP(t *testing.T) {
	// Get first IP
	if res, err := cc.GetIP(cc.IpIDs[0]); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
	// Get second IP
	if res, err := cc.GetIP(cc.IpIDs[1]); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n\n", res)
	}
}

func TestCloudComb_DeleteIP(t *testing.T) {
	// delete first IP
	if err := cc.DeleteIP(cc.IpIDs[0]); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Delete success. \n\n")
	}
	// delete second IP
	if err := cc.DeleteIP(cc.IpIDs[1]); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Print("Delete success. \n\n")
	}
}
