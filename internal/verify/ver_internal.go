package verify

import (
	"encoding/json"
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	baseUrl = "https://verify.twilio.com/v2/"
)

func InternalNewVerificationService(APIClient app.Client, friendlyName string, opts utils.VerOpts)(*ResponseVerifyService, error){

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

	return &r, nil
}

func InternalCompleteVerification(APIClient app.Client, serviceSid, to, code string)(*ResponseConfirmVerification, error){

	requestUrl := baseUrl + "/Services" + serviceSid + "/VerificationCheck"
	method := "POST"

	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("Code", code)
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