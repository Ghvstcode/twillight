package account

//ResponseAccount is the response returned when a User requests Account Information.
type ResponseAccount struct {
	FirstPageURI    string      `json:"first_page_uri"`
	End             int         `json:"end"`
	PreviousPageURI interface{} `json:"previous_page_uri"`
	URI             string      `json:"uri"`
	PageSize        int         `json:"page_size"`
	Start           int         `json:"start"`
	Accounts        []struct {
		Status          string `json:"status"`
		DateUpdated     string `json:"date_updated"`
		AuthToken       string `json:"auth_token"`
		FriendlyName    string `json:"friendly_name"`
		OwnerAccountSid string `json:"owner_account_sid"`
		URI             string `json:"uri"`
		Sid             string `json:"sid"`
		DateCreated     string `json:"date_created"`
		Type            string `json:"type"`
		SubresourceUris struct {
			Addresses             string `json:"addresses"`
			Conferences           string `json:"conferences"`
			SigningKeys           string `json:"signing_keys"`
			Transcriptions        string `json:"transcriptions"`
			ConnectApps           string `json:"connect_apps"`
			Sip                   string `json:"sip"`
			AuthorizedConnectApps string `json:"authorized_connect_apps"`
			Usage                 string `json:"usage"`
			Keys                  string `json:"keys"`
			Applications          string `json:"applications"`
			Recordings            string `json:"recordings"`
			ShortCodes            string `json:"short_codes"`
			Calls                 string `json:"calls"`
			Notifications         string `json:"notifications"`
			IncomingPhoneNumbers  string `json:"incoming_phone_numbers"`
			Queues                string `json:"queues"`
			Messages              string `json:"messages"`
			OutgoingCallerIds     string `json:"outgoing_caller_ids"`
			AvailablePhoneNumbers string `json:"available_phone_numbers"`
			Balance               string `json:"balance"`
		} `json:"subresource_uris"`
	} `json:"accounts"`
	NextPageURI interface{} `json:"next_page_uri"`
	Page        int         `json:"page"`
}
