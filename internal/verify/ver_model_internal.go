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
