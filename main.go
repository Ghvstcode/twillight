package twillight

import (
	"github.com/GhvstCode/twillight/internal/app"
)

type APIClient struct {
	Client app.Client
}

//NewClient creates an authenticated client that can be used to interact with API's
func NewClient(accountSid, authToken string) *APIClient {
	client := app.NewDefaultClient(accountSid, authToken)
	apiClient := APIClient{
		Client: client,
	}

	return &apiClient
}
