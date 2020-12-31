package twillight

import (
	"github.com/GhvstCode/twillight/internal/utils"
	"github.com/GhvstCode/twillight/internal/verify"
)

type VerOptions func(opts *utils.VerOpts)

//OptCodeLength -- Specify the length of the verification code.
//The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
func OptCodeLength(codeLength string) VerOptions{
	return func(s *utils.VerOpts) {
		s.CodeLength = codeLength
	}
}

//OptEnableLookup is to specify whether to perform a lookup with each verification started and return info about the phone number.
func OptEnableLookup(lookup bool) VerOptions{
	return func(s *utils.VerOpts) {
		s.LookupEnabled = lookup
	}
}

//OptEnablePsd2 specifies whether to pass PSD2 transaction parameters when starting a verification.
func OptEnablePsd2(enablePsd2 bool) VerOptions{
	return func(s *utils.VerOpts) {
		s.Psd2Enabled = enablePsd2
	}
}

//OptEnableDoNotShareWarning specifies whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS
func OptEnableDoNotShareWarning(enableWarning bool) VerOptions{
	return func(s *utils.VerOpts) {
		s.DoNotShareWarningEnabled = enableWarning
	}
}

//OptEnableCustomCode specifies Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers
func OptEnableCustomCode(enableCustomCode bool) VerOptions{
	return func(s *utils.VerOpts) {
		s.CustomCodeEnabled = enableCustomCode
	}
}
//NewVerificationService A Verification Service is the set of common configurations used to create and check verifications. One verification service can be used to send multiple verification tokens.
func (c *APIClient) NewVerificationService(FriendlyName string, opts ...VerOptions) (*verify.ResponseVerifyService, error) {
	o := &utils.VerOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := verify.InternalNewVerificationService(c.Client, FriendlyName, *o)
	return res, err
}

func (c *APIClient) UpdateCodeLength(serviceSid, codeLength string)(*verify.ResponseVerifyService,error) {
	res, err := verify.InternalUpdateCodeLength(c.Client, serviceSid,codeLength)
	return res, err
}

func (c *APIClient) UpdateFriendlyName(serviceSid, friendlyName string)(*verify.ResponseVerifyService,error) {
	res, err := verify.InternalUpdateFriendlyName(c.Client, serviceSid,friendlyName)
	return res, err
}

func (c *APIClient) DeleteService(serviceSid string) error {
	err := verify.InternalDeleteService(c.Client, serviceSid)
	return err
}

func (c *APIClient) FetchService(serviceSid string)(*verify.ResponseVerifyService,error){
	res, err := verify.InternalFetchService(c.Client, serviceSid)
	return res, err
}

//func (c *APIClient) SendVerificationToken(serviceSid, to, channel string)(*verify.ResponseSendToken,error) {
//
//}

//StartPsd2Verification is to verify a transaction. You will start by requesting to send a verification code to the user.
func (c *APIClient) StartPsd2Verification(serviceSid,to, channel,amount, payee string)(*verify.ResponseSendToken,error) {
	res, err := verify.InternalStartPsd2Verification(c.Client, serviceSid, to, channel, amount, payee)
	return res, err
}

func (c *APIClient) CompletePsd2Verification(serviceSid,to, code,amount, payee string)(*verify.ResponseConfirmVerification,error) {
	res, err := verify.InternalCompletePsd2Verification(c.Client, serviceSid, to, code, amount, payee)
	return res, err
}