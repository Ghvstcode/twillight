package twillight

import (
	"github.com/GhvstCode/twillight/internal/sms"
	"github.com/GhvstCode/twillight/internal/utils"
)

type SmsClient interface {
	NewOutgoingMessage(to string, from string, body string,  opts ...SmsOptions) (*sms.ResponseSms, error)
	NewOutgoingMediaMessage(to string, from string, msgbody string, mediaUrl string, opts ...SmsOptions) (*sms.ResponseSms, error)
	NewOutgoingWhatsappMessage(to string, from string, body string,  opts ...SmsOptions) (*sms.ResponseSms, error)
	RetrieveAllMessages() (*sms.ResponseGetAllMessages, error)
	RetrieveAllMessagesMedia(messageSid string) (*sms.ResponseAllMessageMedia, error)
	RetrieveMessage(messageSid string)(*sms.ResponseSms,error)
	SendMessageFeedback(messageSid, outcome string)(*sms.ResponseSendMessageFeedback, error)
	DeleteMessage(messageSid string) (*sms.ResponseSms, error)
	UpdateMessage(messageSid, body string)(*sms.ResponseSms, error)
	DeleteMessageMedia(messageSid, mediaSid string) error
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

//NewOutgoingMessage sends a new SMS message to the numbers provided.
func (c *APIClient) NewOutgoingMessage(to string, from string, body string,  opts ...SmsOptions) (*sms.ResponseSms, error){
	o := &utils.SmsOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := sms.InternalNewOutgoingMessage(c.Client, to, from, body, *o)
	return res, err
}

//NewOutgoingWhatsappMessage sends a new whatsapp message to a registered whatsapp Number. WhatsApp requires that your application implement explicit user opt-ins to deliver messages over WhatsApp.
//https://www.twilio.com/docs/whatsapp/api?code-sample=code-send-a-whatsapp-message-and-specify-a-statuscallback-url&code-language=Node.js&code-sdk-version=3.x#using-twilio-phone-numbers-with-whatsapp
func (c *APIClient) NewOutgoingWhatsappMessage(to string, from string, body string,  opts ...SmsOptions) (*sms.ResponseSms, error){
	o := &utils.SmsOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := sms.InternalNewOutgoingWhatsappMessage(c.Client, to, from, body, *o)
	return res, err
}

//NewOutgoingMediaMessage sends a new MMS
func (c *APIClient) NewOutgoingMediaMessage(to string, from string, msgbody string, mediaUrl string, opts ...SmsOptions) (*sms.ResponseSms, error){
	o := &utils.SmsOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := sms.InternalNewOutgoingMediaMessage(c.Client, to, from, msgbody, mediaUrl, *o)
	return res, err
}

//RetrieveAllMessages retrieves all previously sent message
func (c *APIClient) RetrieveAllMessages() (*sms.ResponseGetAllMessages, error){
	res, err := sms.InternalRetrieveAllMessages(c.Client)
	return res, err
}

//RetrieveAllMessagesMedia Lists all media associated with the Account.
func (c *APIClient) RetrieveAllMessagesMedia(messageSid string) (*sms.ResponseAllMessageMedia, error){
	res, err := sms.InternalRetrieveAllMessagesMedia(c.Client, messageSid)
	return res, err
}

//RetrieveAllMessages retrieves a previously sent message
func (c *APIClient) RetrieveMessage(messageSid string)(*sms.ResponseSms,error){
	res, err := sms.InternalRetrieveAMessage(c.Client, messageSid)
	return res, err
}

//https://www.twilio.com/console/sms/insights/delivery?q=(activeInsightsView:overview,filters:!((field:feedback_outcome,filter_type:EQUALS,values:!(UNCONFIRMED))))
//Message Feedback represents the user-reported outcome of a message. For Message Feedback to be sent, the provide feedback option should be set to true when the message is being sent.
func (c *APIClient) SendMessageFeedback(messageSid, outcome string)(*sms.ResponseSendMessageFeedback, error){
	res, err := sms.InternalSendMessageFeedback(c.Client, messageSid, outcome)
	return res, err
}

//UpdateMessage Updates the body of a Message resource. To redact a message, set the body property to an empty string
//https://www.twilio.com/docs/sms/api/message-resource#update-a-message-resource
func (c *APIClient) UpdateMessage(messageSid, body string)(*sms.ResponseSms, error){
	res, err := sms.InternalUpdateMessage(c.Client, messageSid, body)
	return res, err
}


//DeleteMessage deletes a Message record from your account. Once the record is deleted, it will no longer appear in the API and Account Portal logs! On successful deletion, It returns the deleted message.
//https://www.twilio.com/docs/sms/api/message-resource#delete-a-message-resource
func (c *APIClient) DeleteMessage(messageSid string) (*sms.ResponseSms, error){
	res, err := sms.InternalDeleteMessage(c.Client, messageSid)
	return res, err
}

//DeleteMessageMedia deletes the specified media record from your account. Once the record is deleted, it will no longer appear in the API and Account Portal logs! On successful deletion, It returns the deleted message.
//https://www.twilio.com/docs/sms/api/message-resource#message-media-subresources
func (c *APIClient) DeleteMessageMedia(messageSid, mediaSid string) error{
	err := sms.InternalDeleteMessageMedia(c.Client, messageSid, mediaSid)
	return err
}
