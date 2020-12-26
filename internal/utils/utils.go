package utils

const (
	Base_Url = "https://api.twilio.com/2010-04-01/Accounts"
)

//SmsOpts are the options that can be passed to the NewOutgoingMessage function call when creating new SMS message
type SmsOpts struct {
	StatusCallback string
	ProvideFeedback bool
	ValidityPeriod string
}