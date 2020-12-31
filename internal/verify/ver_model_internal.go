package verify

import "time"

//ResponseVerifyService is the Response model for the verification service creation.
type ResponseVerifyService struct {
	Sid                      string `json:"sid"`
	AccountSid               string `json:"account_sid"`
	FriendlyName             string `json:"friendly_name"`
	CodeLength               int    `json:"code_length"`
	LookupEnabled            bool   `json:"lookup_enabled"`
	Psd2Enabled              bool   `json:"psd2_enabled"`
	SkipSmsToLandlines       bool   `json:"skip_sms_to_landlines"`
	DtmfInputRequired        bool   `json:"dtmf_input_required"`
	TtsName                  string `json:"tts_name"`
	DoNotShareWarningEnabled bool   `json:"do_not_share_warning_enabled"`
	CustomCodeEnabled        bool   `json:"custom_code_enabled"`
	Push                     struct {
		IncludeDate      bool        `json:"include_date"`
		ApnCredentialSid string      `json:"apn_credential_sid"`
		FcmCredentialSid interface{} `json:"fcm_credential_sid"`
	} `json:"push"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	URL         string    `json:"url"`
	Links       struct {
		VerificationChecks      string `json:"verification_checks"`
		Verifications           string `json:"verifications"`
		RateLimits              string `json:"rate_limits"`
		MessagingConfigurations string `json:"messaging_configurations"`
		Entities                string `json:"entities"`
		Webhooks                string `json:"webhooks"`
		AccessTokens            string `json:"access_tokens"`
	} `json:"links"`
}
//ResponseSendToken is returned when a verification token is sent too a user.
type ResponseSendToken struct {
	Sid         string    `json:"sid"`
	ServiceSid  string    `json:"service_sid"`
	AccountSid  string    `json:"account_sid"`
	To          string    `json:"to"`
	Channel     string    `json:"channel"`
	Status      string    `json:"status"`
	Valid       bool      `json:"valid"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Lookup      struct {
		Carrier struct {
			ErrorCode         interface{} `json:"error_code"`
			Name              string      `json:"name"`
			MobileCountryCode string      `json:"mobile_country_code"`
			MobileNetworkCode string      `json:"mobile_network_code"`
			Type              string      `json:"type"`
		} `json:"carrier"`
	} `json:"lookup"`
	Amount           string `json:"amount"`
	Payee            string `json:"payee"`
	SendCodeAttempts []struct {
		Time      time.Time   `json:"time"`
		Channel   string      `json:"channel"`
		ChannelID interface{} `json:"channel_id"`
	} `json:"send_code_attempts"`
	URL string `json:"url"`
}

//ResponseConfirmVerification is returned when a verification token is confirmed Successfully.
type ResponseConfirmVerification struct {
	Sid         string      `json:"sid"`
	ServiceSid  string      `json:"service_sid"`
	AccountSid  string      `json:"account_sid"`
	To          string      `json:"to"`
	Channel     string      `json:"channel"`
	Status      string      `json:"status"`
	Valid       bool        `json:"valid"`
	Amount      interface{} `json:"amount"`
	Payee       interface{} `json:"payee"`
	DateCreated time.Time   `json:"date_created"`
	DateUpdated time.Time   `json:"date_updated"`
}

