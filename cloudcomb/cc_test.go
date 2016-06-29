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

func TestCloudComb_UserToken(t *testing.T) {
	if token, expiresIn, err := cc.UserToken(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get token=%s and expires_in=%d\n", token, expiresIn)
		cc.Token = token
	}
}

func TestCloudComb_ContainersImages(t *testing.T) {
	if res, err := cc.ContainersImages(); err != nil {
		fmt.Println(err)
		t.Errorf("Fail to get response. %v", err)
	} else {
		fmt.Printf("Get response: %s\n", res)
	}

}
