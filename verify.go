package twillight

import (
	"github.com/Ghvstcode/twillight/internal/verify"
)

//StartVerification will send a token to the end user through the specified channel.
//Newly created verifications will show a status of pending. Supported channels are sms, call, and email.
func StartVerification(service verify.InternalVerification, to, channel string)(*verify.ResponseSendToken,error) {
	res, err := service.InternalStartVerification(to, channel)
	return res, err
}

//CompleteVerification will check whether the user-provided token is correct.
func CompleteVerification(service verify.InternalVerification, to, code string)(*verify.ResponseConfirmVerification,error) {
	res, err := service.InternalCompleteVerification(to, code)
	return res, err
}

//StartPsd2Verification is to verify a transaction.
//You will start by requesting to send a verification code to the user.
func StartPsd2Verification(service verify.InternalVerification,to, channel,amount, payee string)(*verify.ResponseSendToken,error) {
	res, err := service.InternalStartPsd2Verification(to, channel, amount, payee)
	return res, err
}

//CompletePsd2Verification will check whether the user-provided token for the transaction is correct.
func CompletePsd2Verification(service verify.InternalVerification,to, code,amount, payee string)(*verify.ResponseConfirmVerification,error) {
	res, err := service.InternalCompletePsd2Verification(to, code, amount, payee)
	return res, err
}


