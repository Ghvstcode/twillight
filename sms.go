package twillight

import (
	"github.com/GhvstCode/twillight/internal/sms"
	"net/url"
	"time"
)

type SmsClient interface {
	NewOutgoingMessage(to string, From string, opts SmsOpts) *sms.ResponseSms
	DeleteMessage(messageSid string) sms.ErrorResponse
	RetrieveAllMessages() (*sms.ResponseGetAllMessages, sms.ErrorResponse)
}
type SmsOpts struct {
	StatusCallback url.URL
	MediaUrl url.URL
	provideFeedback bool
	ValidityPeriod time.Duration

}
func (c *APIClient) NewOutgoingMessage(to string, From string, opts SmsOpts) (*sms.ResponseSms, sms.ErrorResponse){
	_ = c

}

//RetrieveAllMessages retrieves all previously sent message
func (c *APIClient) RetrieveAllMessages() (*sms.ResponseGetAllMessages, sms.ErrorResponse){

}

func (c *APIClient) RetrieveAllMessagesMedia() (*sms.ResponseAllMessageMedia, sms.ErrorResponse){

}

//RetrieveAllMessages retrieves a previously sent message
func (c *APIClient) RetrieveMessage(messageSid string){

}

//https://www.twilio.com/console/sms/insights/delivery?q=(activeInsightsView:overview,filters:!((field:feedback_outcome,filter_type:EQUALS,values:!(UNCONFIRMED))))
//Message Feedback represents the user-reported outcome of a message. For Message Feedback to be sent, the provide feedback option should be set to true when the message is being sent.
func (c *APIClient) SendMessageFeedback(messageSid string){

}

//https://www.twilio.com/docs/sms/api/message-resource#update-a-message-resource
func (c *APIClient) UpdateMessage(messageSid string){

}

//https://www.twilio.com/docs/sms/api/message-resource#delete-a-message-resource
func (c *APIClient) DeleteMessage(messageSid string) sms.ErrorResponse{

}