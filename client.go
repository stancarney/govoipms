package govoipms

import (
	"github.com/stancarney/govoipms/v1"
)

func NewV1Client(url, username, password string, debug bool) *v1.Client {
	return &v1.Client{url, username, password, debug}
}