package twillight

import (
	"github.com/Ghvstcode/twillight/internal/service"
	"github.com/Ghvstcode/twillight/internal/utils"
	"github.com/Ghvstcode/twillight/internal/verify"
)

type ServiceOptions func(opts *utils.ServiceOpts)

//OptCodeLength -- Specify the length of the verification code.
//The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
func OptCodeLength(codeLength string) ServiceOptions{
	return func(s *utils.ServiceOpts) {
		s.CodeLength = codeLength
	}
}

//OptEnableLookup is to specify whether to perform a lookup with each verification started and return info about the phone number.
func OptEnableLookup(lookup bool) ServiceOptions{
	return func(s *utils.ServiceOpts) {
		s.LookupEnabled = lookup
	}
}

//OptEnablePsd2 specifies whether to pass PSD2 transaction parameters when starting a verification.
func OptEnablePsd2(enablePsd2 bool) ServiceOptions{
	return func(s *utils.ServiceOpts) {
		s.Psd2Enabled = enablePsd2
	}
}

//OptEnableDoNotShareWarning specifies whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS
func OptEnableDoNotShareWarning(enableWarning bool) ServiceOptions{
	return func(s *utils.ServiceOpts) {
		s.DoNotShareWarningEnabled = enableWarning
	}
}

//OptEnableCustomCode specifies Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers
func OptEnableCustomCode(enableCustomCode bool) ServiceOptions{
	return func(s *utils.ServiceOpts) {
		s.CustomCodeEnabled = enableCustomCode
	}
}


//NewVerificationService Creates a New Verification service which can be used to interact with twilio's verify API's through this library.
//A Verification Service is the set of common configurations used to create and check verifications.
//One verification service can be used to send multiple verification tokens.
//To create a PSD2 enabled service, add the OptEnablePsd2 option when creating a new service.
func (a *Auth) NewVerificationService(FriendlyName string, opts ...ServiceOptions) (*verify.ResponseVerifyService, error) {
	o := &utils.ServiceOpts{}
	for _, opt := range opts {
		opt(o)
	}
	res, err := service.InternalNewVerificationService(a.Client, FriendlyName, *o)
	return res, err
}

//UpdateCodeLength updates the length of the verification code Sent by a service.
func (a *Auth) UpdateCodeLength(serviceSid, codeLength string)(*verify.ResponseVerifyService,error) {
	res, err := service.InternalUpdateCodeLength(a.Client, serviceSid,codeLength)
	return res, err
}

//UpdateFriendlyName Updates the "FriendlyName" of the service in question.
func (a *Auth) UpdateFriendlyName(serviceSid, friendlyName string)(*verify.ResponseVerifyService,error) {
	res, err := service.InternalUpdateFriendlyName(a.Client, serviceSid,friendlyName)
	return res, err
}

//DeleteService Deletes the specified Service
func (a *Auth) DeleteService(serviceSid string) error {
	err := service.InternalDeleteService(a.Client, serviceSid)
	return err
}

//FetchService Fetches a specific service.
func (a *Auth) FetchService(serviceSid string)(*verify.ResponseVerifyService,error){
	res, err := service.InternalRetrieveService(a.Client, serviceSid)
	return res, err
}
