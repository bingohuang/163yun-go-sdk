package cloudcomb

import "fmt"

const (
  Version = "0.0.1"
)

const (
	// defaultEndPoint: access url, support https
	defaultEndPoint = "https://open.c.163.com"

	// defaultConnectTimeout: connection timeout when connect to cloudcomb endpoint
	defaultConnectTimeout = 60
)

// User Agent
func makeUserAgent() string {
	return fmt.Sprintf("CloudComb Go SDK %s", Version)
}

