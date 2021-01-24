package twillight_test

import (
	"errors"
	"github.com/GhvstCode/twillight"
	"github.com/GhvstCode/twillight/internal/sms"
	"github.com/GhvstCode/twillight/internal/utils"
	"reflect"
	"testing"
)

type MockSmsService struct {
	Err error
	Sid string
	AccountSID string
}

func(m *MockSmsService) InternalNewOutgoingMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*sms.ResponseSms, error) {
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to complete verification")
	}

	if to == "" {
		return nil, errors.New("invalid TO number")
	}

	if from == "" {
		return nil, errors.New("invalid FROM number")
	}

	return &sms.ResponseSms{
		AccountSid: m.AccountSID,
		Status: "sent",
		Body: msgbody,
		From: from,
		ErrorCode: nil,
		To: to,
		Sid: m.Sid,
		URI: "/2010-04-01/Accounts/" + m.AccountSID + "/Messages/" + m.Sid+ ".json",
	}, nil
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
		m MockSmsService
		to string
		from string
		body string
		ExpectedErr error
		ExpectedStatus string
		ExpectedURL string
		ExpectedSid string
		ExpectedAccountSid string
		ExpectedErrorCode string
	}{
		{
			m: MockSmsService{
				Err: nil,
				Sid: "12345",
				AccountSID: "4444444",
			},
			to: "987654321",
			from: "000009999777755",
			body: "Hello World",
			ExpectedSid: "12345",
			ExpectedErr: nil,
			ExpectedStatus: "sent",
			ExpectedErrorCode: nil,
			ExpectedURL: "/2010-04-01/Accounts/4444444/Messages/12345.json",

		},
	}

	for _, c := range cases {
		res, err := twillight.NewOutgoingMessage(&c.m, c.to, c.from, c.body)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}
		if res != nil {
			if c.ExpectedSid != res.Sid {
				t.Fatalf("Expected SID to be %s but got %s", c.ExpectedSid, res.Sid)
			}

			if c.ExpectedAccountSid != res.AccountSid {
				t.Fatalf("Expected AccountSID to be %s but got %s", c.ExpectedAccountSid, res.AccountSid)
			}

			if c.ExpectedErrorCode != res.ErrorCode {
				t.Fatalf("Expected ErrorCode to be %s but got %s", c.ExpectedErrorCode, res.ErrorCode)
			}

			if c.ExpectedURL != res.URI {
				t.Fatalf("Expected URL to be %s but got %s", c.ExpectedURL, res.URI)
			}
		}
	}

}