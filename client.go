package crawler

import (
	"github.com/levigross/grequests"
	"time"
)

var client *grequests.Session

// RequestConfig ...
type RequestConfig grequests.RequestOptions

// New ...
func New(userConf RequestConfig) *grequests.Session {
	config := grequests.RequestOptions(userConf)
	config.DialTimeout = 5 * time.Second

	client = grequests.NewSession(&config)
	return client
}
