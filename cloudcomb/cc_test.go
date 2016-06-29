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

func TestToken(t *testing.T) {
	if res, err := cc.UserToken(); err != nil {
		fmt.Println(err)
		t.Errorf("failt to get token. %v", err)
	} else {
		fmt.Printf("Get token: %s\n", res)
	}
}
