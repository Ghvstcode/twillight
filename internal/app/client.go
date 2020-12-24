package app

import (
	"encoding/base64"
	"net/http"
	"net/url"
)

//Config is the result provided when the correct AccountSid & AuthToken are provided
type Config struct {
	HTTPClient *http.Client
}

//Client is the result provided when the correct AccountSid & AuthToken are provided
type Client struct {
	BasicAuth string
	BaseUrl *url.URL
	Configuration Config
}

//basicAuth generates the Basic Auth that is used for interacting with the twillo API
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func NewDefaultClient(username, password string) Client{
	b := basicAuth(username, password)
	c := Client{
		BasicAuth:     b,
		Configuration: struct{ HTTPClient *http.Client }{HTTPClient: http.DefaultClient},
	}

	return c
}