package verify

import (
	"encoding/json"
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	baseUrl = "https://verify.twilio.com/v2/"
)

type InternalVerification interface {
	InternalCompleteVerification(to, code string)(*ResponseConfirmVerification, error)
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

func InternalNewVerificationService( APIClient app.Client, friendlyName string, opts utils.VerOpts)(*ResponseVerifyService, error){
	requestUrl := baseUrl + "/Services"
	method := "POST"

	cl := opts.CodeLength
	if cl == "" {
		cl = "4"
	}

	le := strconv.FormatBool(opts.LookupEnabled)
	cce := strconv.FormatBool(opts.CustomCodeEnabled)
	dnswe := strconv.FormatBool(opts.DoNotShareWarningEnabled)
	pe := strconv.FormatBool(opts.Psd2Enabled)
	Data := url.Values{}
	Data.Set("FriendlyName",friendlyName)
	Data.Set("CodeLength",cl)
	Data.Set("LookupEnabled",le)
	Data.Set("CustomCodeEnabled",cce)
	Data.Set("DoNotShareWarningEnabled",dnswe)
	Data.Set("Psd2Enabled",pe)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseVerifyService
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

	r.Client = APIClient
	return &r, nil
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

func InternalStartVerification(APIClient app.Client, serviceSid, to, channel string)(*ResponseSendToken, error){

	requestUrl := baseUrl + "/Services" + serviceSid + "/Verifications"
	method := "POST"

	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("Channel",channel)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
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

func InternalStartPsd2Verification(APIClient app.Client, serviceSid, to, channel, amount, payee string)(*ResponseSendToken, error){

	requestUrl := baseUrl + "/Services" + serviceSid + "/Verifications"
	method := "POST"

	Data := url.Values{}
	Data.Set("Amount",amount)
	Data.Set("Payee",payee)
	Data.Set("To",to)
	Data.Set("Channel",channel)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
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

func InternalCompletePsd2Verification(APIClient app.Client, serviceSid, to, code, amount, payee string)(*ResponseConfirmVerification, error){

	requestUrl := baseUrl + "/Services/" + serviceSid + "/VerificationCheck"
	method := "POST"

	Data := url.Values{}
	Data.Set("Amount",amount)
	Data.Set("Payee",payee)
	Data.Set("To",to)
	Data.Set("Code",code)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
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

func InternalUpdateCodeLength(APIClient app.Client, serviceSid,codeLength string)(*ResponseVerifyService, error){

	requestUrl := baseUrl + "/Services/" + serviceSid
	method := "POST"

	Data := url.Values{}
	Data.Set("CodeLength",codeLength)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseVerifyService
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

func InternalUpdateFriendlyName(APIClient app.Client, serviceSid,friendlyName string)(*ResponseVerifyService, error){

	requestUrl := baseUrl + "/Services/" + serviceSid
	method := "POST"

	Data := url.Values{}
	Data.Set("FriendlyName",friendlyName)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseVerifyService
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

func InternalDeleteService(APIClient app.Client, serviceSid string) error {

	requestUrl := baseUrl + "/Services/" + serviceSid
	method := "DELETE"


	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)

	res, err := client.Do(req)

	if err != nil {
		return &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	if res.StatusCode  != http.StatusNoContent {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return &e

	}


	return &e
}

func InternalFetchService(APIClient app.Client, serviceSid string) (*ResponseVerifyService, error) {

	requestUrl := baseUrl + "/Services/" + serviceSid
	method := "GET"

	client := APIClient.Configuration.HTTPClient

	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseVerifyService
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