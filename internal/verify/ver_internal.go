package verify

import (
	"encoding/json"
	"github.com/GhvstCode/twillight/internal/app"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	baseUrl = "https://verify.twilio.com/v2/"
)

type InternalVerification interface {
	InternalCompleteVerification(to, code string)(*ResponseConfirmVerification, error)
	InternalStartVerification(to, channel string)(*ResponseSendToken, error)
	InternalStartPsd2Verification(to, channel, amount, payee string)(*ResponseSendToken, error)
	InternalCompletePsd2Verification(to, code, amount, payee string)(*ResponseConfirmVerification, error)
}

type ResponseVerifyService struct {
	app.Client
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


func (s *ResponseVerifyService) InternalCompleteVerification(to, code string)(*ResponseConfirmVerification, error){

	requestUrl := baseUrl + "/Services" + s.Sid + "/VerificationCheck"
	method := "POST"

	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("Code", code)
	DataReader := strings.NewReader(Data.Encode())

	client := s.Client.Configuration.HTTPClient
	//client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", s.Client.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseConfirmVerification
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e
	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func (s *ResponseVerifyService) InternalStartVerification(to, channel string)(*ResponseSendToken, error){

	requestUrl := baseUrl + "/Services" + s.Sid + "/Verifications"
	method := "POST"

	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("Channel",channel)
	DataReader := strings.NewReader(Data.Encode())

	client := s.Client.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", s.Client.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSendToken
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func (s *ResponseVerifyService)InternalStartPsd2Verification(to, channel, amount, payee string)(*ResponseSendToken, error){

	requestUrl := baseUrl + "/Services" + s.Sid + "/Verifications"
	method := "POST"

	Data := url.Values{}
	Data.Set("Amount",amount)
	Data.Set("Payee",payee)
	Data.Set("To",to)
	Data.Set("Channel",channel)
	DataReader := strings.NewReader(Data.Encode())

	client :=s.Client.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", s.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSendToken
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func (s *ResponseVerifyService)InternalCompletePsd2Verification(to, code, amount, payee string)(*ResponseConfirmVerification, error){

	requestUrl := baseUrl + "/Services/" + s.Sid + "/VerificationCheck"
	method := "POST"

	Data := url.Values{}
	Data.Set("Amount",amount)
	Data.Set("Payee",payee)
	Data.Set("To",to)
	Data.Set("Code",code)
	DataReader := strings.NewReader(Data.Encode())

	client := s.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", s.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseConfirmVerification
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

