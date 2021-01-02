package main

import (
	"fmt"
	"github.com/GhvstCode/twillight"
)

func main() {

	c := twillight.NewClient("ACd4859955d9ff9fb86b0a6daabd2bd699", "f7c9a17f472979d2841c0d7a5e7495c6")
	res, err := c.NewOutgoingMessage("+2347032541112", "+16592045850", "Hello Tobi", twillight.OptProvideFeedback(true))

	fmt.Println("Authorization", c.Client.BasicAuth)

	if err != nil {
		fmt.Print("An Error Occurred: ", err)
	}

	fmt.Println("Message Sent: ", res)

}

