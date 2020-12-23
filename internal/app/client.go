package app

import (
	"encoding/base64"
	"net/http"
)

////Client is the result provided when the correct AccountSid & AuthToken are provided
//type Config struct {
//	BasicAuth string
//	HTTPClient *http.Client
//}

//Client is the result provided when the correct AccountSid & AuthToken are provided
type Client struct {
	BasicAuth string
	HTTPClient *http.Client
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}