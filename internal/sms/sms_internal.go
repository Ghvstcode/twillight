package sms

import (
	"encoding/json"
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)
type InternalSMSInterface interface {
	InternalNewOutgoingMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*ResponseSms, error)
	InternalNewOutgoingWhatsappMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*ResponseSms, error)
	InternalNewOutgoingMediaMessage(to string, from string, msgbody string, mediaurl string, opts utils.SmsOpts) (*ResponseSms, error)
	InternalRetrieveAllMessagesMedia(messageSid string) (*ResponseAllMessageMedia, error)
	InternalRetrieveAllMessages() (*ResponseGetAllMessages, error)
	InternalRetrieveAMessage(MessageSid string) (*ResponseSms, error)
	InternalUpdateMessage(MessageSid, body string) (*ResponseSms, error)
	InternalSendMessageFeedback(MessageSid, Outcome string) (*ResponseSendMessageFeedback, error)
	InternalDeleteMessage(MessageSid string)  error
	InternalDeleteMessageMedia(MessageSid string, MediaSid string) error
}

type MessageClient struct {
	Tc app.InternalAuth
}

func (m *MessageClient)InternalNewOutgoingMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*ResponseSms, error){

		requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages.json"
		method := "POST"

	vp := opts.ValidityPeriod
	if vp == "" {
		vp = "14400"
	}
	pf := strconv.FormatBool(opts.ProvideFeedback)
	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("From",from)
	Data.Set("Body",msgbody)
	Data.Set("ValidityPeriod",vp)
	Data.Set("ProvideFeedback",pf)
	Data.Set("StatusCallback",opts.StatusCallback)
	DataReader := strings.NewReader(Data.Encode())

		//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

		client := m.Tc.Configuration.HTTPClient
		//Errors from the API request usually have a
		req, _ := http.NewRequest(method, requestUrl, DataReader)

		req.BasicAuth()

		req.Header.Add("Authorization", m.Tc.BasicAuth)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)

		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}

		defer res.Body.Close()

	var e app.ErrorResponse
		var r ResponseSms
	if res.StatusCode  != http.StatusCreated {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			//fmt.Print("INTERNAL_SMS_MARSHALL_ERR", res.StatusCode)
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

func (m *MessageClient)InternalNewOutgoingWhatsappMessage(to string, from string, msgbody string, opts utils.SmsOpts) (*ResponseSms, error){

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages.json"
	method := "POST"

	vp := opts.ValidityPeriod
	if vp == "" {
		vp = "14400"
	}
	pf := strconv.FormatBool(opts.ProvideFeedback)
	Data := url.Values{}
	Data.Set("To","whatsapp:" + to)
	Data.Set("From","whatsapp:" + from)
	Data.Set("Body",msgbody)
	Data.Set("ValidityPeriod",vp)
	Data.Set("ProvideFeedback",pf)
	Data.Set("StatusCallback",opts.StatusCallback)
	DataReader := strings.NewReader(Data.Encode())

	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
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

func (m *MessageClient)InternalNewOutgoingMediaMessage(to string, from string, msgbody string, mediaurl string, opts utils.SmsOpts) (*ResponseSms, error){

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages.json"
	method := "POST"

	vp := opts.ValidityPeriod
	if vp == "" {
		vp = "14400"
	}
	pf := strconv.FormatBool(opts.ProvideFeedback)
	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("From",from)
	Data.Set("Body",msgbody)
	Data.Set("ValidityPeriod",vp)
	Data.Set("ProvideFeedback",pf)
	Data.Set("MediaUrl",mediaurl)
	Data.Set("StatusCallback",opts.StatusCallback)
	DataReader := strings.NewReader(Data.Encode())

	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)


	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
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

func (m *MessageClient)InternalRetrieveAllMessagesMedia(messageSid string) (*ResponseAllMessageMedia, error){

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages" + messageSid+ "/Media.json"
	method := "GET"

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseAllMessageMedia
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

func (m *MessageClient)InternalRetrieveAllMessages() (*ResponseGetAllMessages, error){

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages.json"
	method := "GET"

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseGetAllMessages
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

func (m *MessageClient)InternalRetrieveAMessage(MessageSid string) (*ResponseSms, error){

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages"+ MessageSid +".json"
	method := "GET"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
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

func (m *MessageClient)InternalUpdateMessage(MessageSid, body string) (*ResponseSms, error){

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages/"+ MessageSid +".json"
	method := "POST"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	Data := url.Values{}
	Data.Set("Body",body)
	DataReader := strings.NewReader(Data.Encode())

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
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

func (m *MessageClient)InternalSendMessageFeedback(MessageSid, Outcome string) (*ResponseSendMessageFeedback, error){

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages/"+ MessageSid +"/Feedback.json"
	method := "POST"

	Data := url.Values{}
	Data.Set("Outcome",Outcome)
	DataReader := strings.NewReader(Data.Encode())

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSendMessageFeedback
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

func (m *MessageClient)InternalDeleteMessage(MessageSid string) error {

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages/"+ MessageSid +".json"
	method := "DELETE"



	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)
	res, err := client.Do(req)

	if err != nil {
		return &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	if res.StatusCode  != http.StatusNoContent{
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return &e

	}
	return nil
}

func (m *MessageClient)InternalDeleteMessageMedia(MessageSid string, MediaSid string) error {

	requestUrl := m.Tc.BaseUrl + "/Accounts/" + m.Tc.AccountSid + "/Messages/"+ MessageSid +"/Media" + MediaSid + ".json"
	method := "DELETE"

	client := m.Tc.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", m.Tc.BasicAuth)
	res, err := client.Do(req)

	if err != nil {
		return  &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	if res.StatusCode  != http.StatusNoContent{
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return &e

	}
	return nil

	//return e.Error()

}
