package sms

//ResponseSms is the response gotten when a message is successfully sent
type ResponseSms struct {
	Sid                 string      `json:"sid"`
	DateCreated         string      `json:"date_created"`
	DateUpdated         string      `json:"date_updated"`
	DateSent            interface{} `json:"date_sent"`
	AccountSid          string      `json:"account_sid"`
	To                  string      `json:"to"`
	From                string      `json:"from"`
	MessagingServiceSid interface{} `json:"messaging_service_sid"`
	Body                string      `json:"body"`
	Status              string      `json:"status"`
	NumSegments         string      `json:"num_segments"`
	NumMedia            string      `json:"num_media"`
	Direction           string      `json:"direction"`
	APIVersion          string      `json:"api_version"`
	Price               interface{} `json:"price"`
	PriceUnit           string      `json:"price_unit"`
	ErrorCode           interface{} `json:"error_code"`
	ErrorMessage        interface{} `json:"error_message"`
	URI                 string      `json:"uri"`
	SubresourceUris     struct {
		Media string `json:"media"`
	} `json:"subresource_uris"`
}

//ResponseGetAllMessages is the response gotten when A user requests for all the SMS messages sent from the account
type ResponseGetAllMessages struct {
	FirstPageURI    string      `json:"first_page_uri"`
	End             int         `json:"end"`
	PreviousPageURI interface{} `json:"previous_page_uri"`
	Messages []ResponseSms
	URI         string      `json:"uri"`
	PageSize    int         `json:"page_size"`
	Start       int         `json:"start"`
	NextPageURI interface{} `json:"next_page_uri"`
	Page        int         `json:"page"`
}

//ResponseSendMessageFeedback is the response gotten when a a user sends a feedback to confirm or unconfirm an SMS message!
type ResponseSendMessageFeedback struct {
	AccountSid  string `json:"account_sid"`
	MessageSid  string `json:"message_sid"`
	Outcome     string `json:"outcome"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	URI         string `json:"uri"`
}


type MediaList struct {
	Sid         string `json:"sid"`
	AccountSid  string `json:"account_sid"`
	ParentSid   string `json:"parent_sid"`
	ContentType string `json:"content_type"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	URI         string `json:"uri"`
}
type ResponseAllMessageMedia struct {
	FirstPageURI    string        `json:"first_page_uri"`
	End             int           `json:"end"`
	MediaList       []MediaList `json:"media_list"`
	PreviousPageURI interface{}   `json:"previous_page_uri"`
	URI             string        `json:"uri"`
	PageSize        int           `json:"page_size"`
	Start           int           `json:"start"`
	NextPageURI     interface{}   `json:"next_page_uri"`
	Page            int           `json:"page"`
}