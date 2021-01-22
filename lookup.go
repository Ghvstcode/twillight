package twillight

import (
	"github.com/GhvstCode/twillight/internal/lookup"
	"github.com/GhvstCode/twillight/internal/utils"
)

type LookupAddon func(opts *utils.LookupAddons)

//NomoroboSpamScoreAddon Detect Robocallers with Lookup and Nomorobo Spam Score Add-on
func NomoroboSpamScoreAddon() LookupAddon{
	return func(s *utils.LookupAddons) {
		s.Addon = "nomorobo_spamscore"
	}
}
//PayfoneAddon  is a Deterministic TCPA Compliance with Lookup and Payfone Add-on
func PayfoneAddon() LookupAddon{
	return func(s *utils.LookupAddons) {
		s.Addon = "payfone_tcpa_compliance"
	}
}

//Get additional caller information with Lookup and Whitepages Pro Caller ID Add-on
func WhitePagesAddon() LookupAddon{
	return func(s *utils.LookupAddons) {
		s.Addon = "whitepages_pro_caller_id"
	}
}

//NewLookupClient creates a new Lookup function which can be used with the functions associated with phone number loookup
func (a *Auth)NewLookupClient() *lookup.ClientLookup{
	return &lookup.ClientLookup{
		Cl: a.Client,
	}
}

//LookupPhoneNumber Is used to get information about a specified phone number.
//https://www.twilio.com/docs/lookup/api
//Lookup also supports Twilio Add-ons, enabling you to retrieve information from a multitude of 3rd party data sources, available via the Twilio Marketplace.
//You can add Lookup-supported add-ons by visiting the Twilio console to enabling the add-on, making sure you have 'Lookups' selected!
func LookupPhoneNumber(LookupClient lookup.InternalLookup, phoneNumber string, addons ...LookupAddon)(*lookup.ResponseLookup, error){
	o := &utils.LookupAddons{}
	for _, opt := range addons {
		opt(o)
	}
	res, err := LookupClient.InternalNewPhoneLookup(phoneNumber, *o)
	return res, err
}
