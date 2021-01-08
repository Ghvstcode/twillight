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


	return &verify.ResponseSendToken{
		Sid: m.Sid,
		To: to,
		Valid: true,
		Channel: channel,
	}, nil
}

func (m *MockVerifyService)InternalStartPsd2Verification(to, channel, amount, payee string)(*verify.ResponseSendToken, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to start verification")
	}

	if to == "" {
		return nil, errors.New("invalid TO number")
	}

	if amount == "" {
		return nil, errors.New("invalid Amount")
	}

	if payee == "" {
		return nil, errors.New("invalid Payee")
	}

	if len(channel) < 3{
		return nil, errors.New("invalid Channel")
	}

	return &verify.ResponseSendToken{
		Sid: m.Sid,
		To: to,
		Valid: true,
	}, nil
}

func (m *MockVerifyService)InternalCompletePsd2Verification(to, channel, amount, payee string)(*verify.ResponseConfirmVerification, error){
	if m.Err != nil {
		return nil, errors.New("an Error Occurred, Unable to start verification")
	}

	if to == "" {
		return nil, errors.New("invalid TO number")
	}

	if amount == "" {
		return nil, errors.New("invalid Amount")
	}

	if payee == "" {
		return nil, errors.New("invalid Payee")
	}

	if len(channel) < 3{
		return nil, errors.New("invalid Channel")
	}

	return &verify.ResponseConfirmVerification{
		Sid: m.Sid,
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
		channel string
		ExpectedSid string
		ExpectedErr error
		ExpectedValid bool
		ExpectedURL string
	}{
		{
			m: MockVerifyService{
				Err: nil,
				CodeLength: 4,
				Sid: "VA12345",
			},
			to: "987654321",
			channel: "email",
			ExpectedSid: "VA12345",
			ExpectedErr: nil,
			ExpectedValid: true,
			ExpectedURL: "https://verify.twilio.com/v2/Services/VA12345",

		},
	}

	for _, c := range cases {
		res, err := twillight.StartVerification(&c.m, c.to, c.channel)
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

func TestStartPsd2Verification(t *testing.T) {
	cases := [] struct{
		m MockVerifyService
		to string
		channel string
		amount string
		payee string
		ExpectedSid string
		ExpectedErr error
		ExpectedValid bool
		ExpectedURL string
	}{
		{
			m: MockVerifyService{
				Err: nil,
				CodeLength: 4,
				Sid: "VA12345",
			},
			to: "987654321",
			channel: "email",
			amount: "",
			payee: "",
			ExpectedSid: "VA12345",
			ExpectedErr: errors.New("invalid Amount"),
		},
		{
			m: MockVerifyService{
				Err: nil,
				CodeLength: 4,
				Sid: "VA12345",
			},
			to: "987654321",
			channel: "email",
			amount: "123",
			payee: "456",
			ExpectedSid: "VA12345",
			ExpectedErr: nil,
			ExpectedURL: "https://verify.twilio.com/v2/Services/VA12345",
			ExpectedValid: true,
		},
	}

	for _, c := range cases {
		res, err := twillight.StartPsd2Verification(&c.m, c.to, c.channel, c.amount, c.payee)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Fatalf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if res != nil{
			//t.Skip("Skipped the rest of the tests")
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
}

func TestCompletePsd2Verification(t *testing.T) {
	cases := [] struct{
		m MockVerifyService
		to string
		code string
		amount string
		payee string
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
				Sid: "VA12345",
			},
			to: "987654321",
			code: "9876",
			amount: "546",
			payee: "444",
			ExpectedSid: "VA12345",
			ExpectedErr: nil,
			ExpectedValid: true,
			ExpectedChannel: "SMS",

		},
	}

	for _, c := range cases {
		res, err := twillight.CompletePsd2Verification(&c.m, c.to, c.code, c.amount, c.payee)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Fatalf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if res != nil{
			//t.Skip("Skipped the rest of the tests")
			if c.ExpectedSid != res.Sid {
				t.Fatalf("Expected SID to be %s but got %s", c.ExpectedSid, res.Sid)
			}

			if c.ExpectedValid != res.Valid {
				t.Fatalf("Expected Valid field to be %t but got %t", c.ExpectedValid, res.Valid)
			}
		}


	}
}