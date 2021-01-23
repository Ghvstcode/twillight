package twillight_test

import (
	"github.com/GhvstCode/twillight"
	"github.com/GhvstCode/twillight/internal/sms"
	"github.com/GhvstCode/twillight/internal/utils"
	"reflect"
	"testing"
)

type MockSmsService struct {
	Err error
	Sid string
}

func(m *MockSmsService) InternalNewOutgoingMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*sms.ResponseSms, error) {

}
func(m *MockSmsService)InternalNewOutgoingWhatsappMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*sms.ResponseSms, error){

}
func(m *MockSmsService)InternalNewOutgoingMediaMessage(to string, from string, msgbody string, mediaurl string, opts utils.SmsOpts) (*sms.ResponseSms, error){

}
func(m *MockSmsService)InternalRetrieveAllMessagesMedia(messageSid string) (*sms.ResponseAllMessageMedia, error){

}
func(m *MockSmsService)InternalRetrieveAllMessages() (*sms.ResponseGetAllMessages, error){

}
func(m *MockSmsService)InternalRetrieveAMessage(MessageSid string) (*sms.ResponseSms, error){

}
func(m *MockSmsService)InternalUpdateMessage(MessageSid, body string) (*sms.ResponseSms, error){

}
func(m *MockSmsService)InternalSendMessageFeedback(MessageSid, Outcome string) (*sms.ResponseSendMessageFeedback, error){

}
func(m *MockSmsService)InternalDeleteMessage(MessageSid string) (*sms.ResponseSms, error){

}
func(m *MockSmsService)InternalDeleteMessageMedia(MessageSid string, MediaSid string) error {

}

func TestNewOutgoingMessage(t *testing.T) {
	cases := [] struct{
		m MockVerifyService
		to string
		code string
		ExpectedSid string
		ExpectedErr error
		ExpectedValid bool
		ExpectedChannel string
	}{
		{
			m: MockVerifyService{
				VerifyChannel: "SMS",
				Err: nil,
				CodeLength: 4,
				Sid: "12345",
			},
			to: "987654321",
			code: "9876",
			ExpectedSid: "12345",
			ExpectedErr: nil,
			ExpectedValid: true,
			ExpectedChannel: "SMS",

		},
	}

	for _, c := range cases {
		res, err := twillight.CompleteVerification(&c.m, c.to, c.code)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if c.ExpectedSid != res.Sid {
			t.Fatalf("Expected SID to be %s but got %s", c.ExpectedSid, res.Sid)
		}

		if c.ExpectedValid != res.Valid {
			t.Fatalf("Expected Valid field to be %t but got %t", c.ExpectedValid, res.Valid)
		}

		if c.ExpectedChannel != res.Channel {
			t.Fatalf("Expected Valid field to be %s but got %s", c.ExpectedChannel, res.Channel)
		}
	}

}