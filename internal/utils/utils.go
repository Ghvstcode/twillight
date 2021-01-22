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

//VerOpts are the options that can be passed when creating new Verify Service
type VerOpts struct {
	//The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
	CodeLength string
	//Whether to perform a lookup with each verification started and return info about the phone number.
	LookupEnabled bool
	//Whether to pass PSD2 transaction parameters when starting a verification.
	Psd2Enabled bool
	//Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS
	DoNotShareWarningEnabled bool
	//Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
	CustomCodeEnabled bool
}

type LookupAddons struct {
	Addon string
}