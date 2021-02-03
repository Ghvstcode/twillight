package twillight_test

import (
	"errors"
	"github.com/Ghvstcode/twillight"
	"github.com/Ghvstcode/twillight/internal/sms"
	"github.com/Ghvstcode/twillight/internal/utils"
	"reflect"
	"testing"
)

type MockSmsService struct {
	Err error
	Sid string
	AccountSID string
}

//YIKEEESSS! I SHOULD PROBABLY SEPARATE THESE INTERFACES
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
		ErrorCode: "0",
		To: to,
		Sid: m.Sid,
		URI: "/2010-04-01/Accounts/" + m.AccountSID + "/Messages/" + m.Sid+ ".json",
	}, nil
}
func(m *MockSmsService)InternalNewOutgoingWhatsappMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*sms.ResponseSms, error){
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
		ErrorCode: "0",
		To: to,
		Sid: m.Sid,
		URI: "/2010-04-01/Accounts/" + m.AccountSID + "/Messages/" + m.Sid+ ".json",
	}, nil
}
func(m *MockSmsService)InternalNewOutgoingMediaMessage(to string, from string, msgbody string, mediaurl string, opts utils.SmsOpts) (*sms.ResponseSms, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to complete verification")
	}

	if mediaurl == "" {
		return nil, errors.New("no Media URL")
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
		ErrorCode: "0",
		To: to,
		Sid: m.Sid,
		URI: "/2010-04-01/Accounts/" + m.AccountSID + "/Messages/" + m.Sid+ ".json",
	}, nil
}
func(m *MockSmsService)InternalRetrieveAllMessagesMedia(messageSid string) (*sms.ResponseAllMessageMedia, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to Retrieve Message Media")
	}

	if messageSid == "" {
		return nil, errors.New("no Media URL")
	}


	return &sms.ResponseAllMessageMedia{
		URI:             "",
	}, nil
}
func(m *MockSmsService)InternalRetrieveAllMessages() (*sms.ResponseGetAllMessages, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to Retrieve Message Media")
	}

	return &sms.ResponseGetAllMessages{
		URI:             "",
	}, nil
}
func(m *MockSmsService)InternalRetrieveAMessage(MessageSid string) (*sms.ResponseSms, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to complete verification")
	}

	if MessageSid == "" {
		return nil, errors.New("no Message SID")
	}

	return &sms.ResponseSms{
		AccountSid: m.AccountSID,
		Status: "sent",
		ErrorCode: nil,
		Sid: MessageSid,
		URI: "/2010-04-01/Accounts/" + m.AccountSID + "/Messages/" + m.Sid+ ".json",
	}, nil
}
func(m *MockSmsService)InternalUpdateMessage(MessageSid, body string) (*sms.ResponseSms, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to complete verification")
	}

	if MessageSid == "" {
		return nil, errors.New("no Message SID")
	}

	if body == "" {
		return nil, errors.New("no message body")
	}

	return &sms.ResponseSms{
		AccountSid: m.AccountSID,
		Status: "sent",
		Body: body,
		ErrorCode: nil,
		Sid: MessageSid,
		URI: "/2010-04-01/Accounts/" + m.AccountSID + "/Messages/" + m.Sid+ ".json",
	}, nil
}
func(m *MockSmsService)InternalSendMessageFeedback(MessageSid, Outcome string) (*sms.ResponseSendMessageFeedback, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to complete verification")
	}

	if MessageSid == "" {
		return nil, errors.New("no Message SID")
	}

	if Outcome == "" {
		return nil, errors.New("no message body")
	}

	return &sms.ResponseSendMessageFeedback{
		AccountSid:  m.AccountSID,
		MessageSid:  MessageSid,
		Outcome:     Outcome,
	}, nil
}
func(m *MockSmsService)InternalDeleteMessage(MessageSid string) error {
	if m.Err != nil {
		return errors.New("an Error Occurred, Unable to complete verification")
	}

	if MessageSid == "" {
		return errors.New("no Message SID")
	}

	return  nil
}
func(m *MockSmsService)InternalDeleteMessageMedia(MessageSid string, MediaSid string) error {
	if m.Err != nil {
		return errors.New("an Error Occurred, Unable to complete verification")
	}

	if MessageSid == "" {
		return errors.New("no Message SID")
	}

	if MediaSid == "" {
		return errors.New("no message body")
	}

	return nil
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
			ExpectedAccountSid: "4444444",
			ExpectedErr: nil,
			ExpectedStatus: "sent",
			ExpectedErrorCode: "0",
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

func TestNewOutgoingWhatsappMessage(t *testing.T) {
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
			ExpectedErrorCode: "0",
			ExpectedAccountSid: "4444444",
			ExpectedURL: "/2010-04-01/Accounts/4444444/Messages/12345.json",

		},
	}

	for _, c := range cases {
		res, err := twillight.NewOutgoingWhatsappMessage(&c.m, c.to, c.from, c.body)
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

func TestNewOutgoingMediaMessage(t *testing.T) {
	cases := [] struct{
		m MockSmsService
		to string
		from string
		body string
		mediaUrl string
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
			ExpectedAccountSid: "4444444",
			ExpectedStatus: "sent",
			mediaUrl: "https://www.mediaURL.com",
			ExpectedErrorCode: "0",
			ExpectedURL: "/2010-04-01/Accounts/4444444/Messages/12345.json",

		},
		{
			m: MockSmsService{
				Err: nil,
				Sid: "12345",
				AccountSID: "4444444",
			},
			to: "",
			from: "000009999777755",
			body: "Hello World",
			mediaUrl: "https://www.mediaURL.com",
			ExpectedSid: "",
			ExpectedAccountSid: "4444444",
			ExpectedErr: errors.New("invalid TO number"),
			ExpectedStatus: "",
			ExpectedErrorCode: "0",
			ExpectedURL: "/2010-04-01/Accounts/4444444/Messages/12345.json",

		},
	}

	for _, c := range cases {
		res, err := twillight.NewOutgoingMediaMessage(&c.m, c.to, c.from, c.body, c.mediaUrl)
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
			if c.ExpectedStatus != res.Status {
				t.Fatalf("Expected AccountSID to be %s but got %s", c.ExpectedStatus, res.Status)
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