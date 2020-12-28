package twillight

import (
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/sms"
	"github.com/GhvstCode/twillight/internal/utils"
	"net/http"
	"net/url"
	"time"
)

type SmsClient interface {
	NewOutgoingMessage(to string, From string, opts ...SmsOptions) *sms.ResponseSms
	DeleteMessage(messageSid string) ErrorResponse
	RetrieveAllMessages() (*sms.ResponseGetAllMessages, ErrorResponse)
}



type SmsOptions func(opts *utils.SmsOpts)

//OptStatusCallback This is the URL Twilio should call using the status_callback_method to send status information to your application. If specified, we POST these message status changes to the URL: queued, failed, sent, delivered, or undelivered
func OptStatusCallback(url string) SmsOptions{
	return func (s *utils.SmsOpts){
		s.StatusCallback = url
	}
}

//OptProvideFeedback Whether to confirm delivery of the message. Set this value to true if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the Message Feedback API. This parameter is false by default.
func OptProvideFeedback(feedback bool) SmsOptions {
	return func (s *utils.SmsOpts){
		s.ProvideFeedback = feedback
	}
}

//OptValidityPeriod How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds.
func OptValidityPeriod(period string) SmsOptions{
	return func (s *utils.SmsOpts){
		s.ValidityPeriod= period
	}
}

func (c *APIClient) NewOutgoingMessage(to string, from string, body string,  opts ...SmsOptions) (*sms.ResponseSms, error){
	o := &utils.SmsOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := sms.InternalNewOutgoingMessage(c.Client, to, from, body, *o)
	return res, err
}

func (c *APIClient) NewOutgoingMediaMessage(to string, From string, mediaUrl string, opts ...SmsOptions) (*sms.ResponseSms, app.ErrorResponse){


}

//RetrieveAllMessages retrieves all previously sent message
func (c *APIClient) RetrieveAllMessages() (*sms.ResponseGetAllMessages, app.ErrorResponse){

}

func (c *APIClient) RetrieveAllMessagesMedia() (*sms.ResponseAllMessageMedia, app.ErrorResponse){

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
func (c *APIClient) DeleteMessage(messageSid string) app.ErrorResponse{

}