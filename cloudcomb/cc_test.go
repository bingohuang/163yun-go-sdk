package cloudcomb

import (
	"fmt"
	"os"
	"testing"
)

// run go test by: CC_APPKEY="your app key" CC_APPSECRET="your app secret" go test -v
var (
	appKey    = os.Getenv("CC_APPKEY")
	appSecret = os.Getenv("CC_APPSECRET")
	cc        = NewCC(appKey, appSecret)
)

/*=== user start 1 ===*/
func TestCloudComb_PostUserToken(t *testing.T) {
	if token, expiresIn, err := cc.PostUserToken(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get token=%s and expires_in=%d\n", token, expiresIn)
		cc.Token = token
	}
}

/*=== containers start 9 ===*/
func TestCloudComb_GetContainersImages(t *testing.T) {
	if res, err := cc.GetContainersImages(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
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
	if res, err := cc.GetContainer("274320"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_GetContainerFlow(t *testing.T) {
	if res, err := cc.GetContainerFlow("274320"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
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
