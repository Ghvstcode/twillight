package twillight

import (
	"github.com/GhvstCode/twillight/internal/app"
)

type APIClient struct {
	Client app.Client
}

//ErrorResponse is the default error response for twilio API's
type ErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}



//NewClient creates an authenticated client that can be used to interact with API's
func NewClient(accountSid, authToken string) *APIClient{
	client := app.NewDefaultClient(accountSid, authToken)
	apiClient := APIClient{
		Client: client,
	}


	return &apiClient
}


