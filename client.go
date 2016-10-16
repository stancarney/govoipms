package govoipms

import (
	"github.com/stancarney/govoipms/v1"
)

func NewV1Client(url, username, password string, debug bool) *v1.VOIPClient {
	return &v1.VOIPClient{url, username, password, debug}
}