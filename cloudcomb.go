package cloudcomb

import "fmt"

const (
  Version = "0.0.1"
)

// User Agent
func makeUserAgent() string {
	return fmt.Sprintf("CloudComb Go SDK %s", Version)
}
