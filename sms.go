package twillight

type test struct {}
type SmsOpts struct {
	callback
	provideFeedback
}
func (c *APIClient) NewOutgoingMessage(to string, From string, ){
	_ = c

}

//RetrieveAllMessages retrieves a previously sent message
func (c *APIClient) RetrieveAllMessages(){

}

//RetrieveAllMessages retrieves a previously sent message
func (c *APIClient) RetrieveMessage(messageSid string){

}
//
func (c *APIClient) ConfirmMessage(messageSid string){

}

func (c *APIClient) UnconfirmMessage(messageSid string){

}