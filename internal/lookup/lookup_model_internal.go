package lookup

type Carrier  struct {
	ErrorCode         interface{} `json:"error_code"`
	MobileCountryCode string      `json:"mobile_country_code"`
	MobileNetworkCode string      `json:"mobile_network_code"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
}

//ResponseLookup is returned by Twilio Lookup API
type ResponseLookup struct {
	CallerName interface{} `json:"caller_name"`
	Carrier   Carrier `json:"carrier"`
	CountryCode    string      `json:"country_code"`
	NationalFormat string      `json:"national_format"`
	PhoneNumber    string      `json:"phone_number"`
	AddOns         interface{} `json:"add_ons"`
	URL            string      `json:"url"`

}