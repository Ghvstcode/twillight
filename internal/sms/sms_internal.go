package sms

import "github.com/GhvstCode/twillight/internal/utils"

type S interface {
	NewOutgoingMessage(to string, From string, opts ...utils.SmsOptions) *sms.ResponseSms
	DeleteMessage(messageSid string) ErrorResponse
	RetrieveAllMessages() (*sms.ResponseGetAllMessages, ErrorResponse)
}
func () InternalNewOutgoingMessage() {

}
//Type A! which should not be imported.