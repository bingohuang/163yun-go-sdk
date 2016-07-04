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
func TestCloudComb_ContainersImages(t *testing.T) {
	if res, err := cc.ContainersImages(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_Containers(t *testing.T) {
	if res, err := cc.Containers(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_Container(t *testing.T) {
	if res, err := cc.Container("274320"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_ContainerFlow(t *testing.T) {
	if res, err := cc.ContainerFlow("274320"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

/*=== clusters(apps) start 8 ===*/
func TestCloudComb_ClustersImages(t *testing.T) {
	if res, err := cc.ClustersImages(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_Clusters(t *testing.T) {
	if res, err := cc.Clusters(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_Cluster(t *testing.T) {
	if res, err := cc.Cluster("413529"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

/*=== repositories start 4 ===*/
func TestCloudComb_Repositories(t *testing.T) {
	if res, err := cc.Repositories(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_Repository(t *testing.T) {
	if res, err := cc.Repository("22103"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

/*=== secret-keys start 4 ===*/
func TestCloudComb_SecretKeys(t *testing.T) {
	if res, err := cc.SecretKeys(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}

func TestCloudComb_SecretKey(t *testing.T) {
	if res, err := cc.SecretKey("196"); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}
}
