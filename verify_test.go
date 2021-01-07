package twillight_test

import (
	"errors"
	"github.com/GhvstCode/twillight"
	"github.com/GhvstCode/twillight/internal/verify"
	"reflect"
	"testing"
)

type MockVerifyService struct {
	//APIClient *twillight.APIClient
	CodeLength int
	Err error
	VerifyChannel string
	Sid string
}

func (m *MockVerifyService) InternalCompleteVerification(to, code string)(*verify.ResponseConfirmVerification, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to complete verification")
	}

	if to == "" {
		return nil, errors.New("invalid TO number")
	}

	if code == "" {
		return nil, errors.New("invalid confirmation code")
	}

	return &verify.ResponseConfirmVerification{
		Sid: m.Sid,
		To: to,
		Valid: true,
		Channel: m.VerifyChannel,
	}, nil
}

func (m *MockVerifyService) InternalStartVerification(to, channel string)(*verify.ResponseSendToken, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to start verification")
	}

	if to == "" {
		return nil, errors.New("invalid TO number")
	}

	if len(channel) < 3{
		return nil, errors.New("invalid Channel")
	}
	if channel != "call"{
		return nil, errors.New("invalid Channel")
	} else if channel != "email"{
		return nil, errors.New("invalid Channel")
	}

	return &verify.ResponseSendToken{
		Sid: m.Sid,
		To: to,
		Valid: true,
		Channel: m.VerifyChannel,
	}, nil
}

func (m *MockVerifyService)InternalStartPsd2Verification(to, channel, amount, payee string)(*verify.ResponseSendToken, error){
	if m.Err != nil {
		return nil, m.Err
	}

	if to == "" {
		return nil, m.Err
	}

	return &verify.ResponseSendToken{
		Sid: "12345",
		To: to,
		Valid: true,
	}, nil
}


func TestCompleteVerification(t *testing.T) {
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

func TestStartVerification(t *testing.T) {
	cases := [] struct{
		m MockVerifyService
		to string
		code string
		ExpectedSid string
		ExpectedErr error
		ExpectedValid bool
		ExpectedURL string
	}{
		{
			m: MockVerifyService{
				VerifyChannel: "email",
				Err: nil,
				CodeLength: 4,
				Sid: "VA12345",
			},
			to: "987654321",
			code: "9876",
			ExpectedSid: "VA12345",
			ExpectedErr: nil,
			ExpectedValid: true,
			ExpectedURL: "https://verify.twilio.com/v2/Services/VA12345",

		},
	}

	for _, c := range cases {
		res, err := twillight.StartVerification(&c.m, c.to, c.code)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Fatalf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if c.ExpectedSid != res.Sid {
			t.Fatalf("Expected SID to be %s but got %s", c.ExpectedSid, res.Sid)
		}

		if c.ExpectedValid != res.Valid {
			t.Fatalf("Expected Valid field to be %t but got %t", c.ExpectedValid, res.Valid)
		}

		if c.ExpectedURL != "https://verify.twilio.com/v2/Services/" + res.Sid {
			t.Fatalf("Expected Valid field to be %s but got %s", c.ExpectedURL, res.Channel)
		}
	}

}