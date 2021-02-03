package account

import (
	"encoding/json"
	"github.com/Ghvstcode/twillight/internal/app"
	"net/http"
)

func InternalGetAccountInfo(APIClient app.InternalAuth) (*ResponseAccount, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/.json"
	method := "GET"

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseAccount
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