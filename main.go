package twillight

import (
	"github.com/GhvstCode/twillight/internal/app"
	"net/http"
)

type Auth struct {
	Client app.InternalAuth
}



//NewAuth creates an authenticated client that can be used to interact with API's
func NewAuth(accountSid, authToken string) *Auth {
	client := app.NewDefaultAuth(accountSid, authToken)
	apiClient := Auth{
		Client: client,
	}

	return &apiClient
}

func (a *Auth) ConfigureHttp(http *http.Client) *Auth{
	a.Client.Configuration.HTTPClient = http
	return a
}
