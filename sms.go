package twillight

import (
	"github.com/GhvstCode/twillight/internal/sms"
	"github.com/GhvstCode/twillight/internal/utils"
)

type SmsOptions func(opts *utils.SmsOpts)

//OptStatusCallback This is the URL Twilio should call using the status_callback_method to send status information to your application. If specified, we POST these message status changes to the URL: queued, failed, sent, delivered, or undelivered
func OptStatusCallback(url string) SmsOptions {
	return func(s *utils.SmsOpts) {
		s.StatusCallback = url
	}
}

//OptProvideFeedback Whether to confirm delivery of the message. Set this value to true if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the Message Feedback API. This parameter is false by default.
func OptProvideFeedback(feedback bool) SmsOptions {
	return func(s *utils.SmsOpts) {
		s.ProvideFeedback = feedback
	}
}

//OptValidityPeriod How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds.
func OptValidityPeriod(period string) SmsOptions {
	return func(s *utils.SmsOpts) {
		s.ValidityPeriod = period
	}
}

//NewSmsClient Creates a New SMSClient for performing SMS operations with Twilight.
func (a *Auth) NewSmsClient () *sms.MessageClient{
	return &sms.MessageClient{
		Tc: a.Client,
	}
}

//NewOutgoingMessage sends a new SMS message to the numbers provided. You can pass additional options! If you want to send messages while in trial mode, you must first verify your 'To' phone number with Twilio. You can verify your phone number by adding it to your Verified Caller IDs in the console.
//If the body of your message is more than 160 GSM-7 characters (or 70 UCS-2characters), Twilio will send the message as a segmented SMS and charge your account accordingly.
func NewOutgoingMessage(s sms.InternalSMSInterface, to string, from string, body string, opts ...SmsOptions) (*sms.ResponseSms, error) {
	o := &utils.SmsOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := s.InternalNewOutgoingMessage(to, from, body, *o)
	return res, err
}

//NewOutgoingWhatsappMessage sends a new whatsapp message to a registered whatsapp Number. WhatsApp requires that your application implement explicit user opt-ins to deliver messages over WhatsApp.
//https://www.twilio.com/docs/whatsapp/api?code-sample=code-send-a-whatsapp-message-and-specify-a-statuscallback-url&code-language=Node.js&code-sdk-version=3.x#using-twilio-phone-numbers-with-whatsapp
func NewOutgoingWhatsappMessage(s sms.InternalSMSInterface,to string, from string, body string, opts ...SmsOptions) (*sms.ResponseSms, error) {
	o := &utils.SmsOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := s.InternalNewOutgoingWhatsappMessage(to, from, body, *o)
	return res, err
}

//NewOutgoingMediaMessage sends a new MMS.
//While you can send text-only messages almost anywhere on the planet, sending media is currently only available in the US and Canada.
//Sending an MMS message looks a lot like sending an SMS, but it includes a new parameter, MediaUrl.
//This parameter specifies the URL of the media you want to include with your message.Twilio supports .gif, .png, or .jpeg content and will format the image on your recipient's device.
//Messages sent via Twilio can include up to 10 media files that have a total size of up to 5MB. Twilio will resize images as necessary for successful delivery based on carrier specifications. Messages with over 5MB of media will not be accepted.
func NewOutgoingMediaMessage(s sms.InternalSMSInterface,to string, from string, msgbody string, mediaUrl string, opts ...SmsOptions) (*sms.ResponseSms, error) {
	o := &utils.SmsOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := s.InternalNewOutgoingMediaMessage(to, from, msgbody, mediaUrl, *o)

	return res, err
}

//RetrieveAllMessages Returns a list of messages associated with your account.
//If you are using this library, the list includes paging information. you may want to use the provided nextpageuri.
//Using the NextPageURI ensures that your next request picks up where it left off and can prevent you from retrieving duplicate data if you are actively sending or receiving messages.
//When getting the list of all messages, results will be sorted on the DateSent field with the most recent messages appearing first.
func RetrieveAllMessages(s sms.InternalSMSInterface) (*sms.ResponseGetAllMessages, error) {
	res, err := s.InternalRetrieveAllMessages()
	return res, err
}

//RetrieveAllMessagesMedia Lists all media associated with the Account.
func RetrieveAllMessagesMedia(s sms.InternalSMSInterface, messageSid string) (*sms.ResponseAllMessageMedia, error) {
	res, err := s.InternalRetrieveAllMessagesMedia(messageSid)
	return res, err
}

//RetrieveMessage retrieves a single message.
func RetrieveMessage(s sms.InternalSMSInterface,messageSid string) (*sms.ResponseSms, error) {
	res, err := s.InternalRetrieveAMessage(messageSid)
	return res, err
}

//Message Feedback represents the user-reported outcome of a message.
//For Message Feedback to be sent, the provide feedback option should be set to true when the message is being sent.
//https://www.twilio.com/console/sms/insights/delivery?q=(activeInsightsView:overview,filters:!((field:feedback_outcome,filter_type:EQUALS,values:!(UNCONFIRMED))))
func SendMessageFeedback(s sms.InternalSMSInterface,messageSid, outcome string) (*sms.ResponseSendMessageFeedback, error) {
	res, err := s.InternalSendMessageFeedback(messageSid, outcome)
	return res, err
}

//UpdateMessage Updates the body of a Message resource.
//To redact a message, set the body property to an empty string
//https://www.twilio.com/docs/sms/api/message-resource#update-a-message-resource
func UpdateMessage(s sms.InternalSMSInterface,messageSid, body string) (*sms.ResponseSms, error) {
	res, err := s.InternalUpdateMessage( messageSid, body)
	return res, err
}

//DeleteMessage deletes a Message record from your account. Once the record is deleted, it will no longer appear in the API and Account Portal logs!
//On successful deletion, It returns No error.
//https://www.twilio.com/docs/sms/api/message-resource#delete-a-message-resource
func DeleteMessage(s sms.InternalSMSInterface,messageSid string) error {
	 err := s.InternalDeleteMessage( messageSid)
	return  err
}

//DeleteMessageMedia deletes the specified media record from your account. Once the record is deleted, it will no longer appear in the API and Account Portal logs! On successful deletion, It returns NO error - An Error value of NIL.
//https://www.twilio.com/docs/sms/api/message-resource#message-media-subresources
func DeleteMessageMedia(s sms.InternalSMSInterface,messageSid, mediaSid string) error {
	err := s.InternalDeleteMessageMedia(messageSid, mediaSid)
return err
}
