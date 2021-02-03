package lookup

import (
	"encoding/json"
	"github.com/Ghvstcode/twillight/internal/app"
	"github.com/Ghvstcode/twillight/internal/utils"
	"net/http"
)

type InternalLookup interface {
	InternalNewPhoneLookup(phone string, addons utils.LookupAddons)(*ResponseLookup, error)
}

type ClientLookup struct {
	Cl app.InternalAuth
}

func (c *ClientLookup)InternalNewPhoneLookup(phone string, addons utils.LookupAddons)(*ResponseLookup, error){
	//c.Cl.BaseUrl = "https://lookups.twilio.com/v1/PhoneNumbers/" + phone
	requestUrl := c.Cl.BaseUrl + phone

	if addons.Addon == "nomorobo_spamscore" {
		//requestUrl = "https://lookups.twilio.com/v1/PhoneNumbers/" + phone +"?AddOns=nomorobo_spamscore"
		requestUrl = c.Cl.BaseUrl + phone + "?AddOns=nomorobo_spamscore"
	}
	if addons.Addon == "payfone_tcpa_compliance" {
		requestUrl = "https://lookups.twilio.com/v1/PhoneNumbers/"+ phone +"?AddOns=payfone_tcpa_compliance&AddOns.payfone_tcpa_compliance.right_party_contacted_date=20160101&Type=carrier"
	}
	if addons.Addon == "whitepages_pro_caller_id" {
		requestUrl = "https://lookups.twilio.com/v1/PhoneNumbers/" + phone + "?AddOns=whitepages_pro_caller_id"
	}
	method := "GET"

	client := c.Cl.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", c.Cl.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseLookup
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, MoreInfo: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}
