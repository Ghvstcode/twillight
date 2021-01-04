package twillight_test

import (
	"github.com/GhvstCode/twillight"
	"github.com/GhvstCode/twillight/internal/verify"
	"testing"
)

type MockVerifyService struct {
	Err error
}

func (m *MockVerifyService) InternalCompleteVerification(to, code string)(*verify.ResponseConfirmVerification, error){
	if m.Err != nil {
		return nil, m.Err
	}

	if to == "" {
		return nil, m.Err
	}

	return &verify.ResponseConfirmVerification{
		Sid: "12345",
		To: to,
	}, nil
}

func TestCompleteVerification(t *testing.T) {
	r := &MockVerifyService{
		Err: nil,
	}

	expectedMsg := "12345"
	res, err := twillight.CompleteVerification(r, "876", "")
	if err != nil {
		t.Fatalf("Expected err to be nil but it was %s", err)
	}

	if expectedMsg != res.Sid {
		t.Fatalf("Expected %s but got %s", expectedMsg, res.Sid)
	}
}
