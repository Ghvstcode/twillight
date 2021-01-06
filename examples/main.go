package main

import (
	"fmt"
	"github.com/GhvstCode/twillight"
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/sms"
)

func NewSms(c *twillight.APIClient)(*sms.ResponseSms, error){
	res, err := c.NewOutgoingMessage("+2347", "+16592045850", "Hello Tobi", twillight.OptProvideFeedback(true))

	return res,err
}

func main() {
	c := twillight.NewClient("ACd4859955d9ff9fb86b0a6daabd2bd699", "f7c9a17f472979d2841c0d7a5e7495c6")
	res, err := NewSms(c)

	if err != nil {
		if err, ok := err.(*app.ErrorResponse); ok {
			fmt.Println("Error Code:", err.ErrorCode())
		}
		fmt.Print("An Error Occurred: ", err)
	}

	fmt.Println("Message Sent: ", res)







	//r, err := newVerify(c)
	//if err != nil {
	//	fmt.Print("An Error Occurred: ", err)
	//	os.Exit(1)
	//}
	//
	//r, err = startVerification(c,r,"", "")
	//if err != nil {
	//	fmt.Print("An Error Occurred: ", err)
	//	os.Exit(1)
	//}
}

