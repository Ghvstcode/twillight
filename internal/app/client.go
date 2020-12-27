package app

import (
	"encoding/base64"
	"net/http"
)

//Config is the result provided when the correct AccountSid & AuthToken are provided
type Config struct {
	HTTPClient *http.Client
}

//Client is the result provided when the correct AccountSid & AuthToken are provided
type Client struct {
	BasicAuth string
	BaseUrl string
	Configuration Config
	AccountSid string
}


//ErrorResponse is the default error response for twilio API's
type ErrorResponse struct {

	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
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
		BaseUrl: "https://api.twilio.com/2010-04-01",
		AccountSid: username,
		Configuration: struct{ HTTPClient *http.Client }{HTTPClient: http.DefaultClient},
	}

	return c
}