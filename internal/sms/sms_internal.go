package sms

import (
	"fmt"
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func InternalNewOutgoingMessage(APIClient app.Client, to string, from string, msgbody string, opts utils.SmsOpts) (*ResponseSms, app.ErrorResponse){

		requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages.json"
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

		client := APIClient.Configuration.HTTPClient
		req, err := http.NewRequest(method, requestUrl, DataReader)
		req.BasicAuth()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		req.Header.Add("Authorization", "Basic QUNkNDg1OTk1NWQ5ZmY5ZmI4NmIwYTZkYWFiZDJiZDY5OTpmN2M5YTE3ZjQ3Mjk3OWQyODQxYzBkN2E1ZTc0OTVjNg==")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(string(body))
}
//Type A! which should not be imported.