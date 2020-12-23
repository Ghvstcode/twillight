package encode

import "encoding/base64"

const (
	Base_Url = "https://api.twilio.com/2010-04-01/Accounts"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}