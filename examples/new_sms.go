package main

import (
	"fmt"
	"github.com/GhvstCode/twillight"
)

func newSms(c *twillight.APIClient){
	res, err := c.NewOutgoingMessage("+2347", "+16592045850", "Hello Tobi", twillight.OptProvideFeedback(true))


	if err != nil {
		fmt.Print("An Error Occurred: ", err)
	}

	fmt.Println("Message Sent: ", res)
}


