package twillight_test

import (
	"errors"
	"github.com/Ghvstcode/twillight"
	"github.com/Ghvstcode/twillight/internal/lookup"
	"github.com/Ghvstcode/twillight/internal/utils"
	"reflect"
	"testing"
)

type MockLookupService struct {}

func (m *MockLookupService) InternalNewPhoneLookup(phone string, addons utils.LookupAddons)(*lookup.ResponseLookup, error){
	if len(phone) < 11 {
		return nil, errors.New("invalid phone number")
	}

	return &lookup.ResponseLookup{
		CallerName: nil,
		Carrier: lookup.Carrier{},
		CountryCode: "US",
		NationalFormat: "",
		PhoneNumber: phone,
		AddOns: nil,
		URL:"https://lookups.twilio.com/v1/PhoneNumbers/"+phone,
	}, nil
}

func TestLookupPhoneNumber(t *testing.T) {
	cases := [] struct{
		m MockLookupService
		phone string
		ExpectedUrl string
		ExpectedErr error
		ExpectedPhone string
	}{
		{
			m: MockLookupService{},
			phone: "+2348987666444443",
			ExpectedUrl:"https://lookups.twilio.com/v1/PhoneNumbers/+2348987666444443",
			ExpectedErr: nil,
			ExpectedPhone: "+2348987666444443",
		},
	}



	for _, c := range cases {
		res, err := twillight.LookupPhoneNumber(&c.m, c.phone)
		if !reflect.DeepEqual(err, c.ExpectedErr) {
			t.Fatalf("Expected err to be %q but it was %q", c.ExpectedErr, err)
		}

		if res != nil{
			//t.Skip("Skipped the rest of the tests")
			if c.ExpectedUrl != res.URL {
				t.Fatalf("Expected URL to be %s but got %s", c.ExpectedUrl, res.URL)
			}

			if c.ExpectedPhone != res.PhoneNumber {
				t.Fatalf("Expected Phone number to be %s but got %s", c.ExpectedPhone, res.PhoneNumber)
			}
		}


	}
}